package controllers

import (
	"context"
	"errors"
	"regexp"

	pb "github.com/malarcon-79/microservices-lab/grpc-protos-go/system/custody"
	"github.com/malarcon-79/microservices-lab/orm-go/dao"
	"github.com/malarcon-79/microservices-lab/orm-go/model"
	"go.uber.org/zap"
)

// Controlador de servicio gRPC
type CustodyServiceController struct {
	logger *zap.SugaredLogger // Logger
	re     *regexp.Regexp     // Expresión regular para validar formato de períodos YYYY-MM
}

// Método a nivel de package, permite construir una instancia correcta del controlador de servicio gRPC
func NewCustodyServiceController() (CustodyServiceController, error) {
	_logger, _ := zap.NewProduction() // Generamos instancia de logger
	logger := _logger.Sugar()

	re, err := regexp.Compile(`^\d{4}\-(0?[1-9]|1[012])$`) // Expresión regular para validar períodos YYYY-MM
	if err != nil {
		return CustodyServiceController{}, err
	}

	instance := CustodyServiceController{
		logger: logger, // Asignamos el logger
		re:     re,     // Asignamos el RegExp precompilado
	}
	return instance, nil // Devolvemos la nueva instancia de este Struct y un puntero nulo para el error
}

func (c *CustodyServiceController) AddCustodyStock(ctx context.Context, msg *pb.CustodyAdd) (*pb.Empty, error) {
	orm := dao.DB.Model(&model.Custody{})

	// TODO: Validaciones

	// Creamos el modelo de datos para almacenamiento
	custody := &model.Custody{
		Period:   msg.Period,
		ClientId: msg.ClientId,
		Stock:    msg.Stock,
		Quantity: int32(msg.Quantity),
	}

	// Insert
	if err := orm.Save(custody).Error; err != nil {
		print(err)
		c.logger.Error("No se pudo agregar la custodia correctamente", err)
		return nil, errors.New("error al guardar")
	}

	// Implementar este método
	return &pb.Empty{}, nil
	// return nil, errors.New("no implementado")
}

func (c *CustodyServiceController) ClosePeriod(ctx context.Context, msg *pb.CloseFilters) (*pb.Empty, error) {
	return nil, errors.New("no implementado")
}

func (c *CustodyServiceController) GetCustody(ctx context.Context, msg *pb.CustodyFilter) (*pb.Custodies, error) {
	// Implementar este método
	return nil, errors.New("no implementado")
}
