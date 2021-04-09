// GENERATED CODE - DO NOT EDIT
// This file provides a way of creating URL's based on all the actions
// found in all the controllers.
package routes

import "github.com/revel/revel"


type tCAssessment struct {}
var CAssessment tCAssessment


func (_ tCAssessment) Init(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("CAssessment.Init", args).URL
}

func (_ tCAssessment) GetAll(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("CAssessment.GetAll", args).URL
}

func (_ tCAssessment) GetByID(
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("CAssessment.GetByID", args).URL
}

func (_ tCAssessment) Create(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("CAssessment.Create", args).URL
}

func (_ tCAssessment) Update(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("CAssessment.Update", args).URL
}

func (_ tCAssessment) Delete(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("CAssessment.Delete", args).URL
}

func (_ tCAssessment) GetEmployees(
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("CAssessment.GetEmployees", args).URL
}

func (_ tCAssessment) DeleteEmployees(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("CAssessment.DeleteEmployees", args).URL
}

func (_ tCAssessment) GetCandidates(
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("CAssessment.GetCandidates", args).URL
}

func (_ tCAssessment) DeleteCandidates(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("CAssessment.DeleteCandidates", args).URL
}


type tCCandidate struct {}
var CCandidate tCCandidate


func (_ tCCandidate) Init(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("CCandidate.Init", args).URL
}

func (_ tCCandidate) GetAll(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("CCandidate.GetAll", args).URL
}

func (_ tCCandidate) GetByID(
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("CCandidate.GetByID", args).URL
}

func (_ tCCandidate) Create(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("CCandidate.Create", args).URL
}

func (_ tCCandidate) Update(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("CCandidate.Update", args).URL
}

func (_ tCCandidate) Delete(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("CCandidate.Delete", args).URL
}


type tCEmployee struct {}
var CEmployee tCEmployee


func (_ tCEmployee) Init(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("CEmployee.Init", args).URL
}

func (_ tCEmployee) GetAll(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("CEmployee.GetAll", args).URL
}

func (_ tCEmployee) GetByID(
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("CEmployee.GetByID", args).URL
}

func (_ tCEmployee) Create(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("CEmployee.Create", args).URL
}

func (_ tCEmployee) Update(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("CEmployee.Update", args).URL
}

func (_ tCEmployee) Delete(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("CEmployee.Delete", args).URL
}


type tCIndex struct {}
var CIndex tCIndex


func (_ tCIndex) Init(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("CIndex.Init", args).URL
}

func (_ tCIndex) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("CIndex.Index", args).URL
}


type tCPosition struct {}
var CPosition tCPosition


func (_ tCPosition) Init(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("CPosition.Init", args).URL
}

func (_ tCPosition) GetAll(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("CPosition.GetAll", args).URL
}


type tCStateAssessment struct {}
var CStateAssessment tCStateAssessment


func (_ tCStateAssessment) Init(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("CStateAssessment.Init", args).URL
}

func (_ tCStateAssessment) GetAll(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("CStateAssessment.GetAll", args).URL
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).URL
}

func (_ tStatic) ServeDir(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeDir", args).URL
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).URL
}

func (_ tStatic) ServeModuleDir(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModuleDir", args).URL
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).URL
}

func (_ tTestRunner) Suite(
		suite string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	return revel.MainRouter.Reverse("TestRunner.Suite", args).URL
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).URL
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).URL
}


