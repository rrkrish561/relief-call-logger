package Contextualizer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	//"github.com/rrkrish561/relief-call-logger/Caller"
	automIO "github.com/rrkrish561/relief-call-logger/Automl"
	"github.com/rrkrish561/relief-call-logger/Message"

	automl "cloud.google.com/go/automl/apiv1"
	automlpb "google.golang.org/genproto/googleapis/cloud/automl/v1"
)

func Contextualize(w http.ResponseWriter, r *http.Request) {

	callerResponse := Message.Message{}

	if err := json.NewDecoder(r.Body).Decode(&callerResponse); err != nil {
		log.Fatalln(err)
		return
	}

	amRequest := automIO.AutomlRequest{
		ProjectID: "grand-harbor-266914",
		Location:  "us",
		ModelID:   "TEN1465862305081196544",
		Content:   callerResponse.Transcript,
	}

	err := languageEntityExtractionPredict(amRequest)
	if err != nil {
		log.Fatalln(err)
		return
	}
}

func languageEntityExtractionPredict(amRequest automIO.AutomlRequest) error {
	ctx := context.Background()
	client, err := automl.NewPredictionClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	req := &automlpb.PredictRequest{
		Name: fmt.Sprintf("projects/%s/locations/%s/models/%s", amRequest.ProjectID, amRequest.Location, amRequest.ModelID),
		Payload: &automlpb.ExamplePayload{
			Payload: &automlpb.ExamplePayload_TextSnippet{
				TextSnippet: &automlpb.TextSnippet{
					Content:  amRequest.Content,
					MimeType: "text/plain",
				},
			},
		},
	}

	resp, err := client.Predict(ctx, req)
	if err != nil {
		return err
	}

	for _, payload := range resp.GetPayload() {
		fmt.Printf("Text extract entity types: %v\n", payload.GetDisplayName())
		fmt.Printf("Text score: %v\n", payload.GetTextExtraction().GetScore())
		textSegment := payload.GetTextExtraction().GetTextSegment()
		fmt.Printf("Text extract entity content: %v\n", textSegment.GetContent())
		fmt.Printf("Text start offset: %v\n", textSegment.GetStartOffset())
		fmt.Printf("Text end offset: %v\n", textSegment.GetEndOffset())
	}

	return nil
}
