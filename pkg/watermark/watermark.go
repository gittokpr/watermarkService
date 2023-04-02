package watermark

import (
	"context"
	"net/http"

	"github.com/gittokpr/watermark-service/internal"
)

type watermarkService struct{}

func NewService() Service {
	return &watermarkService{}
}

func (w *watermarkService) Get(_ context.Context, filters ...internal.Filter) ([]internal.Document, error) {
	doc := internal.Document {
		Content: "book",
		Title: "Harry potter and the half blood prince",
		Author: "J.K. Rowling",
		Topic: "Fiction and magic",
	}

	return []internal.Document{doc}, nil
}

func (w *watermarkService) Status(_ context.Context, ticketID string) (internal.Status, error) {
	return internal.InProgress, nil
}

func (w *watermarkService) Watermark(_ context.Context, ticketID string, mark string) (int, error) {
	return http.StatusOK, nil
}

func (w *watermarkService) AddDocument(_ context.Context, doc *internal.Document) (string, error) {
	newTicketID := shortuuid.New()
	return newTicketID, nil
}

func (w *watermarkService) ServiceStatus(_ context.Context) (int, error)  {
	logger.Log("Checking the Service health...")
	return http.StatusOK, nil
}

var logger log.Logger 

func init() {
	logger = new log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}