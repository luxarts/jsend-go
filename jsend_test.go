package jsend

import (
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewError_WithCode_OK(t *testing.T) {
	// Given
	msg := "Message"
	err := errors.New("some error")
	code := 1234

	// When
	body := NewError(msg, err, code)

	bodyBytes, marshallErr := json.Marshal(body)
	if marshallErr != nil{
		require.Nil(t, marshallErr)
	}

	// Then
	require.Equal(t, "error", body.Status)
	require.Equal(t, "Message", *body.Message)
	require.Equal(t, 1234, *body.Code)
	require.NotNil(t, body.Data)
	require.Equal(t, "some error", body.Data.(string))
	require.Equal(t, "Message (1234)", body.Error())
	require.Equal(t, `{"status":"error","message":"Message","data":"some error","code":1234}`, string(bodyBytes))
}
func TestNewError_WithoutCode_OK(t *testing.T) {
	// Given
	msg := "Message"
	err := errors.New("some error")

	// When
	body := NewError(msg, err)

	bodyBytes, marshallErr := json.Marshal(body)
	if marshallErr != nil{
		require.Nil(t, marshallErr)
	}

	// Then
	require.Equal(t, "error", body.Status)
	require.Equal(t, "Message", *body.Message)
	require.Nil(t, body.Code)
	require.NotNil(t, body.Data)
	require.Equal(t, "some error", body.Data.(string))
	require.Equal(t, "Message", body.Error())
	require.Equal(t, `{"status":"error","message":"Message","data":"some error"}`, string(bodyBytes))
}
func TestNewFail_OK(t *testing.T) {
	// Given
	data := "some data"

	// When
	body := NewFail(data)
	bodyBytes, marshallErr := json.Marshal(body)
	if marshallErr != nil{
		require.Nil(t, marshallErr)
	}

	// Then
	require.Equal(t, "fail", body.Status)
	require.Equal(t, "some data", body.Data)
	require.Nil(t, body.Message)
	require.Nil(t, body.Code)
	require.Equal(t, "some data", body.Error())
	require.Equal(t, `{"status":"fail","data":"some data"}`, string(bodyBytes))
}
func TestNewSuccess_WithData_OK(t *testing.T) {
	// Given
	data := "some data"

	// When
	body := NewSuccess(data)
	bodyBytes, marshallErr := json.Marshal(body)
	if marshallErr != nil{
		require.Nil(t, marshallErr)
	}

	// Then
	require.Equal(t, "success", body.Status)
	require.Equal(t, "some data", body.Data)
	require.Nil(t, body.Message)
	require.Nil(t, body.Code)
	require.Empty(t, body.Error())
	require.Equal(t, `{"status":"success","data":"some data"}`, string(bodyBytes))
}
func TestNewSuccess_WithoutData_OK(t *testing.T) {
	// Given

	// When
	body := NewSuccess(nil)
	bodyBytes, marshallErr := json.Marshal(body)
	if marshallErr != nil{
		require.Nil(t, marshallErr)
	}

	// Then
	require.Equal(t, "success", body.Status)
	require.Nil(t, body.Data)
	require.Nil(t, body.Message)
	require.Nil(t, body.Code)
	require.Empty(t, body.Error())
	require.Equal(t, `{"status":"success","data":null}`, string(bodyBytes))
}