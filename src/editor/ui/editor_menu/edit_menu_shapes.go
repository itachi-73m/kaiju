package editor_menu

import (
	"kaiju/editor/content/content_history"
	"kaiju/editor/editor_interface"
	"kaiju/engine"
	"kaiju/engine/assets"
	"kaiju/matrix"
	"kaiju/rendering"
	"kaiju/rendering/loaders"
	"log/slog"
)

const (
	cubeGLB       = "editor/meshes/cube.glb"
	coneGLB       = "editor/meshes/cone.glb"
	cylinderGLB   = "editor/meshes/cylinder.glb"
	ico_sphereGLB = "editor/meshes/ico_sphere.glb"
	planeGLB      = "editor/meshes/plane.glb"
	sphereGLB     = "editor/meshes/sphere.glb"
	torusGLB      = "editor/meshes/torus.glb"
)

func createShape(name, glb string, ed editor_interface.Editor, host *engine.Host) {
	res, err := loaders.GLTF(glb, host.AssetDatabase())
	if err != nil {
		slog.Error("failed to load the cube mesh", "error", err.Error())
		return
	} else if !res.IsValid() || len(res.Meshes) != 1 {
		slog.Error("cube mesh data corrupted")
		return
	}
	mat, err := host.MaterialCache().Material(assets.MaterialDefinitionBasic)
	if err != nil {
		slog.Error("failed to load the basic material for shape", "error", err)
		return
	}
	resMesh := res.Meshes[0]
	mesh, ok := host.MeshCache().FindMesh(resMesh.MeshName)
	if !ok {
		mesh = rendering.NewMesh(resMesh.MeshName, resMesh.Verts, resMesh.Indexes)
		host.MeshCache().AddMesh(mesh)
	}
	e := ed.CreateEntity(name)
	sd := rendering.ShaderDataBasic{
		ShaderDataBase: rendering.NewShaderDataBase(),
		Color:          matrix.ColorWhite(),
	}
	drawing := rendering.Drawing{
		Renderer:   host.Window.Renderer,
		Material:   mat,
		Mesh:       mesh,
		ShaderData: &sd,
		Transform:  &e.Transform,
	}
	host.Drawings.AddDrawing(drawing)
	e.EditorBindings.AddDrawing(drawing)
	bvh := resMesh.GenerateBVH(host.Threads())
	bvh.Transform = &e.Transform
	e.EditorBindings.Set("bvh", bvh)
	ed.BVH().Insert(bvh)
	e.OnDestroy.Add(func() { bvh.RemoveNode() })
	ed.History().Add(&content_history.ModelOpen{
		Host:   host,
		Entity: e,
		Editor: ed,
	})
}

func (m *Menu) createCone() {
	createShape("Cone", coneGLB, m.editor, m.container.Host)
}

func (m *Menu) createCube() {
	createShape("Cube", cubeGLB, m.editor, m.container.Host)
}

func (m *Menu) createCylinder() {
	createShape("Cylinder", cylinderGLB, m.editor, m.container.Host)
}

func (m *Menu) createIcoSphere() {
	createShape("Ico Sphere", ico_sphereGLB, m.editor, m.container.Host)
}

func (m *Menu) createPlane() {
	createShape("Plane", planeGLB, m.editor, m.container.Host)
}

func (m *Menu) createSphere() {
	createShape("Sphere", sphereGLB, m.editor, m.container.Host)
}

func (m *Menu) createTorus() {
	createShape("Torus", torusGLB, m.editor, m.container.Host)
}
