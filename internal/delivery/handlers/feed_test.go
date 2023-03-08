package handlers

import (
	"depeche/internal/delivery/middleware"
	"depeche/internal/entities"
	"depeche/pkg/apperror"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var feedTestcases = map[string]struct {
	url       string
	method    string
	body      gin.H
	code      int
	err       error
	batchSize string
}{
	"GetFeed 200 Success": {
		url:    "/api/feed",
		method: http.MethodGet,
		body: gin.H{
			"body": gin.H{
				"posts": []entities.Post{
					{
						SenderName: "VK Образование",
						Text: "Сегодня VK Образование совместно с сервисом «Другое Дело» запустили квест «VK: Крафтим карьеру в IT»" +
							" из восьми мини-игр в формате метавселенной Minecraft.\n\nИгра поможет студентам и школьникам познакомиться" +
							" с представителями IT-профессий в формате популярной игры Minecraft. Ведущий и герои в интерактивном формате" +
							" расскажут ребятам о представителях разных IT-профессий: фронтенд-разработчике, дизайнере интерфейсов," +
							" бэкенд-разработчике, игровом разработчике, QA-специалисте, ML-инженере, UX-исследователе, таргетологе." +
							" А ещё проведут по миру карьерных и образовательных проектов VK, в которых в будущем игроки сами смогут" +
							" поучаствовать.\n\nБаллы, полученные за успешное выполнение всех заданий в игре, можно обменять на ценные" +
							" призы в «Другом Деле»: мерч, подписки, технику, билеты на концерты или поездки по стране.\n\nПройти квест" +
							" можно по ссылке: vk.cc/clIuYy.",
						Date:  time.Date(2023, time.February, 27, 17, 16, 3, 0, time.Local),
						Likes: 500,
						Comments: []entities.Comment{
							{
								ID:      1,
								Sender:  "Alexandr Pushkin",
								Date:    time.Date(2023, time.February, 28, 12, 6, 3, 0, time.Local),
								Text:    "И это прекрасно",
								ReplyTo: 0,
							},

							{
								ID:     1,
								Sender: "Michail Lermontov",
								Date:   time.Date(2023, time.February, 28, 12, 6, 3, 0, time.Local),
								Text: "Уважения заслуживают те люди, которые независимо от ситуации," +
									" времени и места, остаются такими же, какие они есть на самом деле.",
								ReplyTo: 0,
							},
						},
					},
				},
			},
		},
		code:      http.StatusOK,
		batchSize: "1",
	},
	"400 BadRequest": {
		url:       "/api/feed",
		method:    http.MethodGet,
		code:      http.StatusBadRequest,
		err:       apperror.BadRequest,
		batchSize: "aaa",
	},
}

func TestFeedHandler(t *testing.T) {

	for name, test := range feedTestcases {
		t.Run(name, func(t *testing.T) {
			query := fmt.Sprintf("%s?batch_size=%s", test.url, test.batchSize)
			request, err := request(test.method, query, test.body)
			require.NoError(t, err)

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, request)

			require.Equal(t, test.code, rr.Code)

			if test.err != nil {
				body, err := json.Marshal(gin.H{
					"status":  middleware.Errors[test.err].Code,
					"message": middleware.Errors[test.err].Message,
				})
				require.NoError(t, err)
				require.Equal(t, body, rr.Body.Bytes())
			} else {
				body, err := json.Marshal(test.body)
				require.NoError(t, err)
				require.Equal(t, body, rr.Body.Bytes())
			}
		})
	}
}
