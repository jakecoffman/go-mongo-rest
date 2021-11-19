package author

import (
	"github.com/gin-gonic/gin"
	"github.com/wwt/go-mongo-rest/lib/db"
	"github.com/wwt/go-mongo-rest/lib/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func ListAuthors(c *gin.Context) {
	var query authorQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(400, "Bad query: "+err.Error())
		return
	}

	opts := options.Find().SetLimit(query.Limit).SetSkip(query.Skip).SetSort(bson.D{{query.Sort, query.Order}})
	cur, err := db.Authors.Find(c, query, opts)
	if err != nil {
		c.JSON(500, "Error querying db "+err.Error())
		return
	}

	authors := []models.Author{}
	if err := cur.All(c, &authors); err != nil {
		c.JSON(500, "Error decoding authors "+err.Error())
		return
	}

	c.JSON(200, authors)
}

func CreateAuthor(c *gin.Context) {
	var author models.Author
	if err := c.ShouldBind(&author); err != nil {
		c.JSON(400, "Bad payload: "+err.Error())
		return
	}
	author.ID = newObject()
	now := time.Now()
	author.Created = &now
	author.Updated = &now

	_, err := db.Authors.InsertOne(c, author)
	if err != nil {
		c.JSON(500, "Failed inserting author: "+err.Error())
		return
	}

	c.JSON(200, author)
}

func PatchAuthor(c *gin.Context) {
	var author models.Author
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(400, "Bad payload: "+err.Error())
		return
	}
	var err error
	author.ID, err = fromHex(c.Param("id"))
	if err != nil {
		c.JSON(400, "Invalid ID")
		return
	}
	now := time.Now()
	author.Updated = &now

	update := bson.D{{"$set", author}}
	opt := options.FindOneAndUpdate().SetReturnDocument(options.After)
	err = db.Authors.FindOneAndUpdate(c, bson.D{{"_id", author.ID}}, update, opt).Decode(&author)
	if err != nil {
		c.JSON(500, "Failed to decode author: "+err.Error())
		return
	}

	c.JSON(200, author)
}

func DeleteAuthor(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(400, "Invalid ID")
		return
	}

	var author models.Author
	err = db.Authors.FindOneAndDelete(c, bson.D{{"_id", id}}).Decode(&author)
	if err != nil {
		c.JSON(404, "Not found")
		return
	}

	c.JSON(200, author)
}

func newObject() *primitive.ObjectID {
	obj := primitive.NewObjectID()
	return &obj
}

func fromHex(hex string) (*primitive.ObjectID, error) {
	obj, err := primitive.ObjectIDFromHex(hex)
	return &obj, err
}
