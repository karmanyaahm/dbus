package storage

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const STORAGE_PATH = "/tmp/database.db"

func TestInit(t *testing.T) {
	assert := assert.New(t)

	db, err := InitStorage("/dev/notexistingfolder/database.db")
	assert.Error(err)
	assert.Nil(db)

	db, err = InitStorage(STORAGE_PATH)
	assert.NoError(err)
	defer os.Remove(STORAGE_PATH)
	assert.NotNil(db)
}

func TestNewConnectionWithGeneratedToken(t *testing.T) {
	assert := assert.New(t)

	db, err := InitStorage(STORAGE_PATH)
	assert.NoError(err)
	defer os.Remove(STORAGE_PATH)

	appID := "app-1"
	appToken := "apptoken-2"
	endpointSettings := "<endpoint>"

	conn := db.NewConnection(appID, appToken, endpointSettings)
	assert.NotNil(conn)
	// be sure that PublicToken is no given value
	assert.NotEqual("", conn.PublicToken)
	assert.NotEqual(appID, conn.PublicToken)
	assert.NotEqual(appToken, conn.PublicToken)
	assert.NotEqual(endpointSettings, conn.PublicToken)

	// that everythink else is given
	assert.Equal(appID, conn.AppID)
	assert.Equal(appToken, conn.AppToken)
	assert.Equal(endpointSettings, conn.Settings)
}

func TestNewConnectionCollision(t *testing.T) {
	assert := assert.New(t)

	db, err := InitStorage(STORAGE_PATH)
	assert.NoError(err)
	defer os.Remove(STORAGE_PATH)

	appID := "app-1"
	publicToken := "public-token-2"
	endpoint := "<endpoint>"

	// create connection
	conn := db.NewConnectionWithToken(appID, "app-token-1", publicToken, endpoint)
	assert.NotNil(conn)

	// collision is nil
	conn = db.NewConnectionWithToken(appID, "app-token-2", publicToken, endpoint)
	assert.Nil(conn)

}

func TestNewConnectionUpdateEndpoint(t *testing.T) {
	assert := assert.New(t)

	db, err := InitStorage(STORAGE_PATH)
	assert.NoError(err)
	defer os.Remove(STORAGE_PATH)

	appID := "app-1"
	appToken := "apptoken-2"
	publicToken := "public-token-2"
	oldEndpointSettings := "endpoint-1"
	newEndpointSettings := "endpoint-2"

	// create connection
	conn := db.NewConnectionWithToken(appID, appToken, publicToken, oldEndpointSettings)
	assert.NotNil(conn)
	conn = db.getFirst(Connection{AppID: appID, AppToken: appToken})
	assert.NotNil(conn)
	assert.Equal(oldEndpointSettings, conn.Settings)

	// save new endpoint on connection
	db.NewConnectionWithToken(appID, appToken, publicToken, newEndpointSettings)
	assert.NotNil(conn)
	conn = db.getFirst(Connection{AppID: appID, AppToken: appToken})
	assert.NotNil(conn)
	assert.Equal(newEndpointSettings, conn.Settings)

}

func TestGetConnectionbyPublic(t *testing.T) {
	assert := assert.New(t)

	db, err := InitStorage(STORAGE_PATH)
	assert.NoError(err)
	defer os.Remove(STORAGE_PATH)

	publicToken := "public-token-2"

	// create multiple connection
	db.NewConnectionWithToken("appid-1", "apptoken-1", "public-token-1", "<endpoint>")
	db.NewConnectionWithToken("appid-1", "apptoken-2", publicToken, "<endpoint>")
	db.NewConnectionWithToken("appid-1", "apptoken-3", "public-token-3", "<endpoint>")

	// find correct connection by public token
	conn := db.GetConnectionbyPublic(publicToken)
	assert.Equal(publicToken, conn.PublicToken)
}

func TestDeleteConnection(t *testing.T) {
	assert := assert.New(t)

	db, err := InitStorage(STORAGE_PATH)
	assert.NoError(err)
	defer os.Remove(STORAGE_PATH)

	appToken := "apptoken-2"

	// create multiple connection
	db.NewConnectionWithToken("appid-1", "apptoken-1", "public-token-1", "<endpoint>")
	db.NewConnectionWithToken("appid-1", appToken, "public-token-2", "<endpoint>")
	db.NewConnectionWithToken("appid-1", "apptoken-3", "public-token-3", "<endpoint>")

	// find correct connection by app token to delete
	conn, err := db.DeleteConnection(appToken)
	assert.NoError(err)
	assert.Equal(appToken, conn.AppToken)

	// unable to delete connection does not exists anymore
	conn, err = db.DeleteConnection(appToken)
	assert.Error(err)
}
