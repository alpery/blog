package blog

import (
	"encoding/json"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogPost struct {
	BlogId        primitive.ObjectID `json:"blogId" bson:"_id,omitempty"`
	Headline      string             `json:"headline" bson:"headline"`
	FeaturedImage []byte             `json:"featuredImage" bson:"featuredImage"`
	Content       json.RawMessage    `json:"content" bson:"content"`
	CreatedAt     time.Time          `json:"createdAt" bson:"createdAt"`
}
