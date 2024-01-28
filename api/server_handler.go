package api

import (
	"context"
	"fmt"
	"os"

	db "github.com/Streamfair/streamfair_user_svc/db/sqlc"
	"github.com/Streamfair/streamfair_user_svc/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Server serves HTTP requests for the streamfair backend service.
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("acctype", validAccountTypes)
	}

	router.GET("/readiness", server.readinessCheck)

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccountByID)
	router.GET("/accounts/owner/:owner", server.getAccountByOwner)
	router.GET("/accounts/owner", server.handleMissingUsername)
	router.GET("/accounts", server.listAccount)
	router.PUT("/accounts/:id", server.updateAccount)
	router.DELETE("/accounts/:id", server.deleteAccount)

	server.router = router
	return server
}

// StartServer starts a new HTTP server on the specified address.
func (server *Server) StartServer(address string) error {
	if err := InitializeDatabase(server.store); err != nil {
		fmt.Fprintf(os.Stderr, "database: error while initializing database: %v\n", err)
		return err
	}
	return server.router.Run(address)
}

// InitializeDatabase creates the initial fixed entries in the database.
func InitializeDatabase(store db.Store) error {
	accountTypes := util.GetAccountTypeStruct()
	arg := db.ListAccountTypesParams{
		Limit:  int32(len(accountTypes)),
		Offset: 0,
	}
	accountTypesInDB, err := store.ListAccountTypes(context.Background(), arg)
	if err != nil {
		return err
	}

	// Convert accountTypesInDB into a map for faster lookup
	accountTypesMap := make(map[int32]bool)
	for _, accountTypeInDB := range accountTypesInDB {
		accountTypesMap[accountTypeInDB.ID] = true
	}

	var errs []error
	for _, accountType := range accountTypes {
		if !accountTypesMap[int32(accountType.ID)] {
			_, err := store.CreateAccountType(context.Background(), db.CreateAccountTypeParams{
				Type:        accountType.Type,
				Permissions: accountType.Permissions,
				IsArtist:    accountType.IsArtist,
				IsProducer:  accountType.IsProducer,
				IsWriter:    accountType.IsWriter,
				IsLabel:     accountType.IsLabel,
			})
			if err != nil {
				errs = append(errs, err)
			}
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("%q", errs)
	}

	fmt.Println("Database initialized.")
	return nil
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
