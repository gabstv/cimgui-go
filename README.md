# cimgui-go for ebiten-imgui

This is a **fork** of the AllenDang/cimgui-go library with the sole purpose to be compatible with [Ebitengine](https://ebitengine.org/).
The original GLFW 3.3 and OpenGL 3.2 backends were **removed** because the C/C++ dependencies are incompatible (and in this use case, redundant).
You should most definitely use the **original** repository (AllenDang/cimgui-go) if you are not going to use [ebiten-imgui](https://github.com/gabstv/ebiten-imgui/).

## cimgui-go

This project aims to generate go wrapper for Dear ImGui.

### Current solution is:
1. Use cimgui's lua generator to generate function and struct definition as json.
2. Generate proper go code from the definition (via manual crafted go program `/cmd/codegen`).
3. Use the backend implementation from imgui (currently glfw and opengl3).
4. Use github workflow to compile cimgui and glfw to static lib and place them in /lib folder for further link. 

### Naming convention

- For functions, 'Im/ImGui/ig' is trimmed.
- If function comes from `imgui_internal.h`, `Internal` prefix is added.
- Struct fields (if any exported) are prefixed with word `Field`

### Function coverage
Currently most of the functions are generated, except memory related stuff (eg. memory allocator, storage management, etc...).
If you find any function is missing, report an issue.

### Generate binding
Install [GNU make](https://www.gnu.org/software/make/manual/make.html) and run `make` to re-generate bunding.

### Update

To update to the latest version of dependencies, run `make update`.
After doing this, commit changes and navigate to GitHub.
In Actions tab, manually trigger workflows for each platform.

### How does it work?

- `cimgui/` directory holds C binding for C++ Dear ImGui libraries
- generator bases on `cimgui/{package_name}_templates` and generates all necessary GO/C code
- `libs/` contains pre-built shared libraries. `cimgui.go` includes and uses to decrease building time.
