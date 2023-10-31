// Code generated by cmd/codegen from https://github.com/gabstv/cimgui-go.
// DO NOT EDIT.

#pragma once

#include "cimnodes_wrapper.h"

#ifdef __cplusplus
extern "C" {
#endif

extern void wrap_EmulateThreeButtonMouse_SetModifier(EmulateThreeButtonMouse *EmulateThreeButtonMousePtr, const bool* v);
extern const bool* wrap_EmulateThreeButtonMouse_GetModifier(EmulateThreeButtonMouse *self);
extern void wrap_ImNodesIO_SetEmulateThreeButtonMouse(ImNodesIO *ImNodesIOPtr, EmulateThreeButtonMouse v);
extern EmulateThreeButtonMouse wrap_ImNodesIO_GetEmulateThreeButtonMouse(ImNodesIO *self);
extern void wrap_ImNodesIO_SetLinkDetachWithModifierClick(ImNodesIO *ImNodesIOPtr, LinkDetachWithModifierClick v);
extern LinkDetachWithModifierClick wrap_ImNodesIO_GetLinkDetachWithModifierClick(ImNodesIO *self);
extern void wrap_ImNodesIO_SetMultipleSelectModifier(ImNodesIO *ImNodesIOPtr, MultipleSelectModifier v);
extern MultipleSelectModifier wrap_ImNodesIO_GetMultipleSelectModifier(ImNodesIO *self);
extern void wrap_ImNodesIO_SetAltMouseButton(ImNodesIO *ImNodesIOPtr, int v);
extern int wrap_ImNodesIO_GetAltMouseButton(ImNodesIO *self);
extern void wrap_ImNodesIO_SetAutoPanningSpeed(ImNodesIO *ImNodesIOPtr, float v);
extern float wrap_ImNodesIO_GetAutoPanningSpeed(ImNodesIO *self);
extern void wrap_ImNodesStyle_SetGridSpacing(ImNodesStyle *ImNodesStylePtr, float v);
extern float wrap_ImNodesStyle_GetGridSpacing(ImNodesStyle *self);
extern void wrap_ImNodesStyle_SetNodeCornerRounding(ImNodesStyle *ImNodesStylePtr, float v);
extern float wrap_ImNodesStyle_GetNodeCornerRounding(ImNodesStyle *self);
extern void wrap_ImNodesStyle_SetNodePadding(ImNodesStyle *ImNodesStylePtr, ImVec2 v);
extern ImVec2 wrap_ImNodesStyle_GetNodePadding(ImNodesStyle *self);
extern void wrap_ImNodesStyle_SetNodeBorderThickness(ImNodesStyle *ImNodesStylePtr, float v);
extern float wrap_ImNodesStyle_GetNodeBorderThickness(ImNodesStyle *self);
extern void wrap_ImNodesStyle_SetLinkThickness(ImNodesStyle *ImNodesStylePtr, float v);
extern float wrap_ImNodesStyle_GetLinkThickness(ImNodesStyle *self);
extern void wrap_ImNodesStyle_SetLinkLineSegmentsPerLength(ImNodesStyle *ImNodesStylePtr, float v);
extern float wrap_ImNodesStyle_GetLinkLineSegmentsPerLength(ImNodesStyle *self);
extern void wrap_ImNodesStyle_SetLinkHoverDistance(ImNodesStyle *ImNodesStylePtr, float v);
extern float wrap_ImNodesStyle_GetLinkHoverDistance(ImNodesStyle *self);
extern void wrap_ImNodesStyle_SetPinCircleRadius(ImNodesStyle *ImNodesStylePtr, float v);
extern float wrap_ImNodesStyle_GetPinCircleRadius(ImNodesStyle *self);
extern void wrap_ImNodesStyle_SetPinQuadSideLength(ImNodesStyle *ImNodesStylePtr, float v);
extern float wrap_ImNodesStyle_GetPinQuadSideLength(ImNodesStyle *self);
extern void wrap_ImNodesStyle_SetPinTriangleSideLength(ImNodesStyle *ImNodesStylePtr, float v);
extern float wrap_ImNodesStyle_GetPinTriangleSideLength(ImNodesStyle *self);
extern void wrap_ImNodesStyle_SetPinLineThickness(ImNodesStyle *ImNodesStylePtr, float v);
extern float wrap_ImNodesStyle_GetPinLineThickness(ImNodesStyle *self);
extern void wrap_ImNodesStyle_SetPinHoverRadius(ImNodesStyle *ImNodesStylePtr, float v);
extern float wrap_ImNodesStyle_GetPinHoverRadius(ImNodesStyle *self);
extern void wrap_ImNodesStyle_SetPinOffset(ImNodesStyle *ImNodesStylePtr, float v);
extern float wrap_ImNodesStyle_GetPinOffset(ImNodesStyle *self);
extern void wrap_ImNodesStyle_SetMiniMapPadding(ImNodesStyle *ImNodesStylePtr, ImVec2 v);
extern ImVec2 wrap_ImNodesStyle_GetMiniMapPadding(ImNodesStyle *self);
extern void wrap_ImNodesStyle_SetMiniMapOffset(ImNodesStyle *ImNodesStylePtr, ImVec2 v);
extern ImVec2 wrap_ImNodesStyle_GetMiniMapOffset(ImNodesStyle *self);
extern void wrap_ImNodesStyle_SetFlags(ImNodesStyle *ImNodesStylePtr, ImNodesStyleFlags v);
extern ImNodesStyleFlags wrap_ImNodesStyle_GetFlags(ImNodesStyle *self);
extern void wrap_ImNodesStyle_SetColors(ImNodesStyle *ImNodesStylePtr, unsigned int* v);
extern unsigned int* wrap_ImNodesStyle_GetColors(ImNodesStyle *self);
extern void wrap_LinkDetachWithModifierClick_SetModifier(LinkDetachWithModifierClick *LinkDetachWithModifierClickPtr, const bool* v);
extern const bool* wrap_LinkDetachWithModifierClick_GetModifier(LinkDetachWithModifierClick *self);
extern void wrap_MultipleSelectModifier_SetModifier(MultipleSelectModifier *MultipleSelectModifierPtr, const bool* v);
extern const bool* wrap_MultipleSelectModifier_GetModifier(MultipleSelectModifier *self);

#ifdef __cplusplus
}
#endif
