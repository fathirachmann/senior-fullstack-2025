package service

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/qdrant/go-client/qdrant"
)

type VectorService struct {
	Client *qdrant.Client
}

func NewVectorService(client *qdrant.Client) *VectorService {
	return &VectorService{
		Client: client,
	}
}

func (s *VectorService) CreateCollection(collectionName string, vectorSize uint64) error {
	return s.Client.CreateCollection(context.Background(), &qdrant.CreateCollection{
		CollectionName: collectionName,
		VectorsConfig: qdrant.NewVectorsConfig(&qdrant.VectorParams{
			Size:     vectorSize,
			Distance: qdrant.Distance_Cosine,
		}),
	})
}

func (s *VectorService) UpsertArticle(collectionName string, id uint64, vector []float32, title string) error {
	point := &qdrant.PointStruct{
		Id:      qdrant.NewIDNum(id),
		Vectors: qdrant.NewVectors(vector...),
		Payload: qdrant.NewValueMap(map[string]any{"title": title}),
	}

	_, err := s.Client.Upsert(context.Background(), &qdrant.UpsertPoints{
		CollectionName: collectionName,
		Points:         []*qdrant.PointStruct{point},
	})
	return err
}

func (s *VectorService) SearchArticle(collectionName string, vector []float32) (string, error) {
	limit := uint64(1)
	searchResult, err := s.Client.Query(context.Background(), &qdrant.QueryPoints{
		CollectionName: collectionName,
		Query:          qdrant.NewQuery(vector...),
		Limit:          &limit,
		WithPayload:    qdrant.NewWithPayload(true),
	})
	if err != nil {
		return "", err
	}

	if len(searchResult) == 0 {
		return "", errors.New("no results found")
	}

	payload := searchResult[0].Payload
	if title, ok := payload["title"]; ok {
		return title.GetStringValue(), nil
	}
	return "", errors.New("title not found in payload")
}

// BulkInsert implements the concurrent upsert
func (s *VectorService) BulkInsert(collectionName string, vectors [][]float32, payloads []map[string]interface{}) error {
	if len(vectors) != len(payloads) {
		return errors.New("vectors and payloads length mismatch")
	}

	batchSize := 100
	total := len(vectors)
	var wg sync.WaitGroup
	errChan := make(chan error, total/batchSize+1)

	for i := 0; i < total; i += batchSize {
		end := i + batchSize
		if end > total {
			end = total
		}

		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()

			var points []*qdrant.PointStruct
			for j := start; j < end; j++ {
				id := uuid.New().String()
				points = append(points, &qdrant.PointStruct{
					Id:      qdrant.NewID(id),
					Vectors: qdrant.NewVectors(vectors[j]...),
					Payload: qdrant.NewValueMap(payloads[j]),
				})
			}

			_, err := s.Client.Upsert(context.Background(), &qdrant.UpsertPoints{
				CollectionName: collectionName,
				Points:         points,
			})
			if err != nil {
				errChan <- fmt.Errorf("batch %d-%d failed: %w", start, end, err)
			}
		}(i, end)
	}

	wg.Wait()
	close(errChan)

	for err := range errChan {
		return err
	}

	return nil
}
