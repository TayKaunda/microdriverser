package main 
import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	pb "github.com/sikozonpc/ride-sharing/shared/proto/trip"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

