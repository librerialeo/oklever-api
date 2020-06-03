package websocket

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/librerialeo/oklever-api/pkg/utils"
)

// GetTestClassByTeachersID get test class
func GetTestClassByTeachersID(s *Socket, a *Action) {
	if s.user.ID != 0 {
		testClass, err := s.io.service.GetTestClassByTeacherID(s.user.ID)
		if err != nil {
			s.EmitServerError("No se puede obtener la clase muestra", err)
		} else {
			s.Emit("ADD_TEACHER_TEST_CLASS", testClass)
		}
	}
}

// AddTeachersTestClass set teachers test class
func AddTeachersTestClass(s *Socket, a *Action) {
	data, ok := a.Data.(map[string]interface{})
	if ok && s.user.ID != 0 {
		name, nameOk := data["name"].(string)
		video, videoOk := data["video"].(string)
		if nameOk && videoOk {
			ext := strings.Split(strings.Split(video, ";")[0], "/")[1]
			videoname := fmt.Sprintf("test-class-%d-%d.%s", s.user.ID, time.Now().Unix(), ext)
			err := utils.SaveVideoToFile(os.Getenv("TEST_CLASS_DIR"), videoname, video)
			testClass, err := s.io.service.GetTestClassByTeacherID(s.user.ID)
			if err != nil {
				s.EmitServerError("Error al obtener clase muestra", err)
			}
			if testClass == nil {
				testClass, err := s.io.service.AddTeachersTestClass(s.user.ID, name, video)
				if err != nil {
					s.EmitServerError("Error al guardar la información", err)
				} else {
					s.Emit("ADD_TEACHER_TEST_CLASS", testClass)
				}
			} else {
				// hacer el update
			}
		}
	}
}
