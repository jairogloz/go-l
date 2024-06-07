package team

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/jairogloz/go-l/internal/domain"
)

/*
Create es un método de la estructura Service que crea un nuevo equipo en la base de datos.

Parámetros:
  - team: Un puntero a una estructura domain.Team. Esto permite que el método modifique la estructura original del equipo.
  - ctx: Un objeto context.Context. Esto permite que el método pase el contexto a otros métodos que llama.

El método llama al método Insert del campo Repo de la estructura Service, pasando la estructura Team y el contexto como argumentos.

Si el método Insert devuelve un error, el método Create verifica si el error es un error de clave duplicada. Si lo es, registra el error
y devuelve un nuevo AppError con el código y mensaje de error de clave duplicada. Si el error no es un error de clave duplicada, registra el error
y devuelve un nuevo error que envuelve el error original con un mensaje que indica que hubo un error al crear el equipo.

Valores de retorno:
  - err: Un error que será nil si el equipo se creó correctamente. Si hubo un error, será un objeto de error que describe el fallo.

Nota: Dejo team.CreatedAt comentado porque no se está utilizando en el código, ya que aun no se hace el core de la funcionalidad.
*/
func (s *Service) Create(ctx context.Context, team *domain.Team) (err error) {
	//now := time.Now().UTC()
	//team.CreatedAt := &now

	err = s.Repo.Insert(ctx, team)
	if err != nil {
		if errors.Is(err, domain.ErrDuplicateKey) {
			log.Fatalln("Duplicate key error")
			appErr := domain.AppError{
				Code: domain.ErrCodeDuplicateKey,
				Msg:  "error creating team: duplicate key error",
			}
			return appErr
		}
		log.Println(err.Error())
		return fmt.Errorf("error creating team: %w", err)
	}
	return nil
}
