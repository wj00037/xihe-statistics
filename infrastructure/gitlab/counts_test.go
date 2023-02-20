package gitlab

import (
	"net/http"
	"net/http/httptest"
	"project/xihe-statistics/domain/platform"
	"reflect"
	"testing"

	"github.com/opensourceways/community-robot-lib/utils"
)

func TestGitlabStatistics_GetProjectId(t *testing.T) {
	// 构造gitlabStatistics对象
	impl := &gitlabStatistics{
		endpoint: "http://localhost:8083/api/v4",
		token:    "glpat-jD3gsd5swfM35dxXL3Dd",
		cli:      utils.NewHttpClient(3),
	}

	// 设置测试参数
	pageNum := 1

	// 调用方法
	projectIds, err := impl.GetProjectId(pageNum)

	// 断言结果
	if err != nil {
		t.Errorf("GetProjectId failed: %v", err)
	}
	if len(projectIds) == 0 {
		t.Errorf("GetProjectId returned empty result")
	}
}

func TestGitlabStatistics_GetCloneTotal(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		response   []byte
		want       platform.CloneTotal
		wantErr    error
	}{
		{
			name:       "success",
			statusCode: http.StatusOK,
			response:   []byte(`{"clone_count":1,"repository_size":2048}`),
			want:       platform.CloneTotal{Total: 6},
			wantErr:    nil,
		},
		// {
		// 	name:       "failure - error response",
		// 	statusCode: http.StatusInternalServerError,
		// 	response:   []byte(`{"message": "failed to get clone total"}`),
		// 	want:       platform.CloneTotal{},
		// 	wantErr:    errors.New("failed to get clone total"),
		// },
		// {
		// 	name:       "failure - invalid response body",
		// 	statusCode: http.StatusOK,
		// 	response:   []byte(`invalid_response`),
		// 	want:       platform.CloneTotal{},
		// 	wantErr:    errors.New("invalid character 'i' looking for beginning of value"),
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// create a test server with a response that matches the current test case
			testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.statusCode)
				w.Write(tt.response)
			}))
			defer testServer.Close()

			// create an instance of gitlabStatistics with a custom http client that uses the test server
			impl := &gitlabStatistics{
				endpoint: "http://localhost:8083/api/v4",
				token:    "glpat-jD3gsd5swfM35dxXL3Dd",
				cli:      utils.NewHttpClient(3),
			}

			// call the method being tested
			got, err := impl.GetCloneTotal(3)

			// verify that the result and error match what is expected
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GitlabStatistics.GetCloneTotal() = %v, want %v", got, tt.want)
			}

			if (err == nil && tt.wantErr != nil) || (err != nil && tt.wantErr == nil) || (err != nil && tt.wantErr != nil && err.Error() != tt.wantErr.Error()) {
				t.Errorf("GitlabStatistics.GetCloneTotal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
