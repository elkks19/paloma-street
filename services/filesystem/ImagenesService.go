package filesystem

// import (
// 	"encoding/base64"
// 	"os"
// 	"path/filepath"
// 	"proyecto/internal/config/env"
// 	"proyecto/services"
// )

// TODO: IMPLEMENTAR ESTA WEA PLIS
type ImagenesService struct {

}

func NewImagenesService() *ImagenesService {
	return &ImagenesService{}
}

func (s *ImagenesService) Save(b64 string, model any) (string, error) {
	// dec, err := base64.StdEncoding.DecodeString(b64)
	// if err != nil {
	// 	return "", err
	// }
	//
	// var fp string
	// switch f := model.(type) {
	// 	case *services.Producto:
	// 		fp = filepath.Join(env.Get("IMAGE_FOLDER"), f.Negocio.Nombre, "productos", f.Nombre)
	//
	// 	case *services.Negocio:
	// 		fp = filepath.Join(env.Get("IMAGE_FOLDER"), f.Nombre)
	//
	// }
	//
	//
	// // fp := filepath.Join("/imagenes", )
	// err = os.MkdirAll(filepath.Dir(fp), 0755)
	// if err != nil {
	// 	fmt.Println("Error creating directories:", err)
	// 	return
	// }
	//
	return "", nil
}

func (s *ImagenesService) Delete(url string) error {
	return nil
}

func (s *ImagenesService) Get(url string) ([]byte, error) {
	return nil, nil
}
