// Code generated by cmd/codegen from https://github.com/gabstv/cimgui-go.
// DO NOT EDIT.


#include <string.h>
#include "cimnodes_wrapper.h"
#include "cimnodes_structs_accessor.h"

void wrap_EmulateThreeButtonMouse_SetModifier(EmulateThreeButtonMouse *EmulateThreeButtonMousePtr, const bool* v) { EmulateThreeButtonMousePtr->Modifier = v; }
const bool* wrap_EmulateThreeButtonMouse_GetModifier(EmulateThreeButtonMouse *self) { return self->Modifier; }
void wrap_ImNodesIO_SetEmulateThreeButtonMouse(ImNodesIO *ImNodesIOPtr, EmulateThreeButtonMouse v) { ImNodesIOPtr->EmulateThreeButtonMouse = v; }
EmulateThreeButtonMouse wrap_ImNodesIO_GetEmulateThreeButtonMouse(ImNodesIO *self) { return self->EmulateThreeButtonMouse; }
void wrap_ImNodesIO_SetLinkDetachWithModifierClick(ImNodesIO *ImNodesIOPtr, LinkDetachWithModifierClick v) { ImNodesIOPtr->LinkDetachWithModifierClick = v; }
LinkDetachWithModifierClick wrap_ImNodesIO_GetLinkDetachWithModifierClick(ImNodesIO *self) { return self->LinkDetachWithModifierClick; }
void wrap_ImNodesIO_SetMultipleSelectModifier(ImNodesIO *ImNodesIOPtr, MultipleSelectModifier v) { ImNodesIOPtr->MultipleSelectModifier = v; }
MultipleSelectModifier wrap_ImNodesIO_GetMultipleSelectModifier(ImNodesIO *self) { return self->MultipleSelectModifier; }
void wrap_ImNodesIO_SetAltMouseButton(ImNodesIO *ImNodesIOPtr, int v) { ImNodesIOPtr->AltMouseButton = v; }
int wrap_ImNodesIO_GetAltMouseButton(ImNodesIO *self) { return self->AltMouseButton; }
void wrap_ImNodesIO_SetAutoPanningSpeed(ImNodesIO *ImNodesIOPtr, float v) { ImNodesIOPtr->AutoPanningSpeed = v; }
float wrap_ImNodesIO_GetAutoPanningSpeed(ImNodesIO *self) { return self->AutoPanningSpeed; }
void wrap_ImNodesStyle_SetGridSpacing(ImNodesStyle *ImNodesStylePtr, float v) { ImNodesStylePtr->GridSpacing = v; }
float wrap_ImNodesStyle_GetGridSpacing(ImNodesStyle *self) { return self->GridSpacing; }
void wrap_ImNodesStyle_SetNodeCornerRounding(ImNodesStyle *ImNodesStylePtr, float v) { ImNodesStylePtr->NodeCornerRounding = v; }
float wrap_ImNodesStyle_GetNodeCornerRounding(ImNodesStyle *self) { return self->NodeCornerRounding; }
void wrap_ImNodesStyle_SetNodePadding(ImNodesStyle *ImNodesStylePtr, ImVec2 v) { ImNodesStylePtr->NodePadding = v; }
ImVec2 wrap_ImNodesStyle_GetNodePadding(ImNodesStyle *self) { return self->NodePadding; }
void wrap_ImNodesStyle_SetNodeBorderThickness(ImNodesStyle *ImNodesStylePtr, float v) { ImNodesStylePtr->NodeBorderThickness = v; }
float wrap_ImNodesStyle_GetNodeBorderThickness(ImNodesStyle *self) { return self->NodeBorderThickness; }
void wrap_ImNodesStyle_SetLinkThickness(ImNodesStyle *ImNodesStylePtr, float v) { ImNodesStylePtr->LinkThickness = v; }
float wrap_ImNodesStyle_GetLinkThickness(ImNodesStyle *self) { return self->LinkThickness; }
void wrap_ImNodesStyle_SetLinkLineSegmentsPerLength(ImNodesStyle *ImNodesStylePtr, float v) { ImNodesStylePtr->LinkLineSegmentsPerLength = v; }
float wrap_ImNodesStyle_GetLinkLineSegmentsPerLength(ImNodesStyle *self) { return self->LinkLineSegmentsPerLength; }
void wrap_ImNodesStyle_SetLinkHoverDistance(ImNodesStyle *ImNodesStylePtr, float v) { ImNodesStylePtr->LinkHoverDistance = v; }
float wrap_ImNodesStyle_GetLinkHoverDistance(ImNodesStyle *self) { return self->LinkHoverDistance; }
void wrap_ImNodesStyle_SetPinCircleRadius(ImNodesStyle *ImNodesStylePtr, float v) { ImNodesStylePtr->PinCircleRadius = v; }
float wrap_ImNodesStyle_GetPinCircleRadius(ImNodesStyle *self) { return self->PinCircleRadius; }
void wrap_ImNodesStyle_SetPinQuadSideLength(ImNodesStyle *ImNodesStylePtr, float v) { ImNodesStylePtr->PinQuadSideLength = v; }
float wrap_ImNodesStyle_GetPinQuadSideLength(ImNodesStyle *self) { return self->PinQuadSideLength; }
void wrap_ImNodesStyle_SetPinTriangleSideLength(ImNodesStyle *ImNodesStylePtr, float v) { ImNodesStylePtr->PinTriangleSideLength = v; }
float wrap_ImNodesStyle_GetPinTriangleSideLength(ImNodesStyle *self) { return self->PinTriangleSideLength; }
void wrap_ImNodesStyle_SetPinLineThickness(ImNodesStyle *ImNodesStylePtr, float v) { ImNodesStylePtr->PinLineThickness = v; }
float wrap_ImNodesStyle_GetPinLineThickness(ImNodesStyle *self) { return self->PinLineThickness; }
void wrap_ImNodesStyle_SetPinHoverRadius(ImNodesStyle *ImNodesStylePtr, float v) { ImNodesStylePtr->PinHoverRadius = v; }
float wrap_ImNodesStyle_GetPinHoverRadius(ImNodesStyle *self) { return self->PinHoverRadius; }
void wrap_ImNodesStyle_SetPinOffset(ImNodesStyle *ImNodesStylePtr, float v) { ImNodesStylePtr->PinOffset = v; }
float wrap_ImNodesStyle_GetPinOffset(ImNodesStyle *self) { return self->PinOffset; }
void wrap_ImNodesStyle_SetMiniMapPadding(ImNodesStyle *ImNodesStylePtr, ImVec2 v) { ImNodesStylePtr->MiniMapPadding = v; }
ImVec2 wrap_ImNodesStyle_GetMiniMapPadding(ImNodesStyle *self) { return self->MiniMapPadding; }
void wrap_ImNodesStyle_SetMiniMapOffset(ImNodesStyle *ImNodesStylePtr, ImVec2 v) { ImNodesStylePtr->MiniMapOffset = v; }
ImVec2 wrap_ImNodesStyle_GetMiniMapOffset(ImNodesStyle *self) { return self->MiniMapOffset; }
void wrap_ImNodesStyle_SetFlags(ImNodesStyle *ImNodesStylePtr, ImNodesStyleFlags v) { ImNodesStylePtr->Flags = v; }
ImNodesStyleFlags wrap_ImNodesStyle_GetFlags(ImNodesStyle *self) { return self->Flags; }
void wrap_ImNodesStyle_SetColors(ImNodesStyle *ImNodesStylePtr, unsigned int* v) { memcpy(ImNodesStylePtr->Colors, v, sizeof(unsigned int)*29); }
unsigned int* wrap_ImNodesStyle_GetColors(ImNodesStyle *self) { return self->Colors; }
void wrap_LinkDetachWithModifierClick_SetModifier(LinkDetachWithModifierClick *LinkDetachWithModifierClickPtr, const bool* v) { LinkDetachWithModifierClickPtr->Modifier = v; }
const bool* wrap_LinkDetachWithModifierClick_GetModifier(LinkDetachWithModifierClick *self) { return self->Modifier; }
void wrap_MultipleSelectModifier_SetModifier(MultipleSelectModifier *MultipleSelectModifierPtr, const bool* v) { MultipleSelectModifierPtr->Modifier = v; }
const bool* wrap_MultipleSelectModifier_GetModifier(MultipleSelectModifier *self) { return self->Modifier; }
