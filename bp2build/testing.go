package bp2build

import (
	"android/soong/android"

	"github.com/google/blueprint"
)

type nestedProps struct {
	Nested_prop string
}

type customProps struct {
	Bool_prop     bool
	Bool_ptr_prop *bool
	// Ensure that properties tagged `blueprint:mutated` are omitted
	Int_prop         int `blueprint:"mutated"`
	Int64_ptr_prop   *int64
	String_prop      string
	String_ptr_prop  *string
	String_list_prop []string

	Nested_props     nestedProps
	Nested_props_ptr *nestedProps
}

type customModule struct {
	android.ModuleBase

	props customProps
}

// OutputFiles is needed because some instances of this module use dist with a
// tag property which requires the module implements OutputFileProducer.
func (m *customModule) OutputFiles(tag string) (android.Paths, error) {
	return android.PathsForTesting("path" + tag), nil
}

func (m *customModule) GenerateAndroidBuildActions(ctx android.ModuleContext) {
	// nothing for now.
}

func customModuleFactoryBase() android.Module {
	module := &customModule{}
	module.AddProperties(&module.props)
	return module
}

func customModuleFactory() android.Module {
	m := customModuleFactoryBase()
	android.InitAndroidModule(m)
	return m
}

type testProps struct {
	Test_prop struct {
		Test_string_prop string
	}
}

type customTestModule struct {
	android.ModuleBase

	props      customProps
	test_props testProps
}

func (m *customTestModule) GenerateAndroidBuildActions(ctx android.ModuleContext) {
	// nothing for now.
}

func customTestModuleFactoryBase() android.Module {
	m := &customTestModule{}
	m.AddProperties(&m.props)
	m.AddProperties(&m.test_props)
	return m
}

func customTestModuleFactory() android.Module {
	m := customTestModuleFactoryBase()
	android.InitAndroidModule(m)
	return m
}

type customDefaultsModule struct {
	android.ModuleBase
	android.DefaultsModuleBase
}

func customDefaultsModuleFactoryBase() android.DefaultsModule {
	module := &customDefaultsModule{}
	module.AddProperties(&customProps{})
	return module
}

func customDefaultsModuleFactoryBasic() android.Module {
	return customDefaultsModuleFactoryBase()
}

func customDefaultsModuleFactory() android.Module {
	m := customDefaultsModuleFactoryBase()
	android.InitDefaultsModule(m)
	return m
}

type bp2buildBlueprintWrapContext struct {
	bpCtx *blueprint.Context
}

func (ctx *bp2buildBlueprintWrapContext) ModuleName(module blueprint.Module) string {
	return ctx.bpCtx.ModuleName(module)
}

func (ctx *bp2buildBlueprintWrapContext) ModuleDir(module blueprint.Module) string {
	return ctx.bpCtx.ModuleDir(module)
}

func (ctx *bp2buildBlueprintWrapContext) ModuleSubDir(module blueprint.Module) string {
	return ctx.bpCtx.ModuleSubDir(module)
}

func (ctx *bp2buildBlueprintWrapContext) ModuleType(module blueprint.Module) string {
	return ctx.bpCtx.ModuleType(module)
}

func (ctx *bp2buildBlueprintWrapContext) VisitAllModulesBlueprint(visit func(blueprint.Module)) {
	ctx.bpCtx.VisitAllModules(visit)
}

func (ctx *bp2buildBlueprintWrapContext) VisitDirectDeps(module android.Module, visit func(android.Module)) {
	ctx.bpCtx.VisitDirectDeps(module, func(m blueprint.Module) {
		if aModule, ok := m.(android.Module); ok {
			visit(aModule)
		}
	})
}
