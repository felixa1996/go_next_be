package domain_user_usecase

import (
	"bytes"
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"

	domain "github.com/felixa1996/go_next_be/app/domain/user"
	dto "github.com/felixa1996/go_next_be/app/domain/user/dto"
	mocks "github.com/felixa1996/go_next_be/mocks/app/domain/user"
)

func TestUserCreate(t *testing.T) {
	// todo change to table test
	t.Parallel()
	logger := zaptest.NewLogger(t)

	reader := bytes.NewReader([]byte("1111111111111111"))
	uuid.SetRand(reader)
	uuid.SetClockSequence(1)

	user := domain.User{
		Id:     "31313131-3131-4131-b131-313131313131",
		Name:   "test2",
		Author: "test2",
	}

	t.Run("Success", func(y *testing.T) {

		mockRepo := new(mocks.UserRepository)
		mockRepo.On("Create", context.TODO(), user).Return(user, nil)

		a := NewUserUsecase(mockRepo, logger, time.Second*2)
		res, err := a.Create(context.TODO(), dto.UserDtoCreateInput{
			Name:   user.Name,
			Author: user.Author,
		})
		if err != nil {
			// assert.Contains(t, err.Error(), tc.ExpectErrorReponse.Error(), "Should error")
		}

		assert.Equal(t, user, res, "Should success")
	})
}
