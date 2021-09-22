package gin

import (
	ggc "gogochat"
	"github.com/gin-gonic/gin"
	"gogochat/inmemory"
	"net/http"
)

type Server struct {
	router *gin.Engine

	UserService    ggc.UserService
	ChannelService ggc.ChannelService
	MessageService ggc.MessageService
}

func NewServer() *Server {
	router := gin.Default()
	userSvc := inmemory.NewUserService()
	channelSvc := inmemory.NewChannelService()
	msgSvc := inmemory.NewMessageService(channelSvc, userSvc)
	s := &Server{
		router: router,
		UserService: userSvc,
		ChannelService: channelSvc,
		MessageService: msgSvc,
	}

	router.POST("/user", s.createUser)
	router.GET("/user/:name", s.getUser)

	return s
}

func (s *Server) getUser(c *gin.Context) {
	name := c.Params.ByName("name")
	u, err := s.UserService.FindUserByUsername(name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"name": name, "status": "user not found"})
	} else {
		c.JSON(http.StatusOK, gin.H{"user": u})
	}
}

func (s *Server) createUser(c *gin.Context) {
	// Parse JSON
	var json struct {
		Username  string `json:"username"`
	}

	if c.Bind(&json) == nil {
		u := &ggc.User{Username: json.Username}
		if err := s.UserService.CreateUser(u); err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "ok", "user": u})
	}
}

func (s *Server) Start() error {
	if err := s.router.Run(":8080"); err != nil {
		return err
	}

	return nil
}
