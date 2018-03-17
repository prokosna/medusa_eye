package main

import (
	"github.com/fgrosse/goldi"
)

type Repository interface {
	Select() string
	Update(value string)
}

type RepositoryImpl struct {
	text string
}

func NewRepositoryImpl() *RepositoryImpl {
	println("New RepositoryImpl")
	return &RepositoryImpl{text: "default"}
}

func (r *RepositoryImpl) Select() string {
	return r.text
}

func (r *RepositoryImpl) Update(value string) {
	r.text = value
}

type Service interface {
	Execute() string
}

type ServiceImpl struct {
	repository Repository
}

func NewServiceImpl(repository Repository) *ServiceImpl {
	println("New ServiceImpl")
	return &ServiceImpl{repository: repository}
}

func (s *ServiceImpl) Execute() string {
	return s.repository.Select()
}

func main() {
	registry := goldi.NewTypeRegistry()
	config := map[string]interface{}{
		"timeout": 42,
	}
	container := goldi.NewContainer(registry, config)

	container.RegisterType("Repository", NewRepositoryImpl)
	container.RegisterType("Service", NewServiceImpl, "@Repository")

	repository := container.MustGet("Repository").(Repository)
	println(repository.Select())
	repository.Update("changed")
	println(repository.Select())
	service := container.MustGet("Service").(Service)
	println(service.Execute())
	repository.Update("again!")
	println(repository.Select())
	println(service.Execute())
	println("End!")
}
