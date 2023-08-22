// Code generated by cmd/codegen from https://github.com/AllenDang/cimgui-go.
// DO NOT EDIT.

package imgui

// #include <stdlib.h>
// #include <memory.h>
// #include "extra_types.h"
// #include "cimplot_wrapper.h"
import "C"
import "unsafe"

type FormatterTimeData struct {
	// TODO: contains unsupported fields
	data unsafe.Pointer
}

func (self FormatterTimeData) handle() (result *C.Formatter_Time_Data, releaseFn func()) {
	result = (*C.Formatter_Time_Data)(self.data)
	return result, func() {}
}

func (self FormatterTimeData) c() (result C.Formatter_Time_Data, fin func()) {
	resultPtr, finFn := self.handle()
	return *resultPtr, finFn
}

func newFormatterTimeDataFromC(cvalue *C.Formatter_Time_Data) *FormatterTimeData {
	result := new(FormatterTimeData)
	result.data = unsafe.Pointer(cvalue)
	return result
}

type PlotAlignmentData struct {
	FieldVertical bool
	FieldPadA     float32
	FieldPadB     float32
	FieldPadAMax  float32
	FieldPadBMax  float32
}

func (self PlotAlignmentData) handle() (result *C.ImPlotAlignmentData, releaseFn func()) {
	result = new(C.ImPlotAlignmentData)
	FieldVertical := self.FieldVertical

	result.Vertical = C.bool(FieldVertical)
	FieldPadA := self.FieldPadA

	result.PadA = C.float(FieldPadA)
	FieldPadB := self.FieldPadB

	result.PadB = C.float(FieldPadB)
	FieldPadAMax := self.FieldPadAMax

	result.PadAMax = C.float(FieldPadAMax)
	FieldPadBMax := self.FieldPadBMax

	result.PadBMax = C.float(FieldPadBMax)
	releaseFn = func() {
	}
	return result, releaseFn
}

func (self PlotAlignmentData) c() (result C.ImPlotAlignmentData, fin func()) {
	resultPtr, finFn := self.handle()
	return *resultPtr, finFn
}

func newPlotAlignmentDataFromC(cvalue *C.ImPlotAlignmentData) *PlotAlignmentData {
	result := new(PlotAlignmentData)
	result.FieldVertical = cvalue.Vertical == C.bool(true)
	result.FieldPadA = float32(cvalue.PadA)
	result.FieldPadB = float32(cvalue.PadB)
	result.FieldPadAMax = float32(cvalue.PadAMax)
	result.FieldPadBMax = float32(cvalue.PadBMax)
	return result
}

type PlotAnnotation struct {
	FieldPos        Vec2
	FieldOffset     Vec2
	FieldColorBg    uint32
	FieldColorFg    uint32
	FieldTextOffset int32
	FieldClamp      bool
}

func (self PlotAnnotation) handle() (result *C.ImPlotAnnotation, releaseFn func()) {
	result = new(C.ImPlotAnnotation)
	FieldPos := self.FieldPos

	result.Pos = FieldPos.toC()
	FieldOffset := self.FieldOffset

	result.Offset = FieldOffset.toC()
	FieldColorBg := self.FieldColorBg

	result.ColorBg = C.ImU32(FieldColorBg)
	FieldColorFg := self.FieldColorFg

	result.ColorFg = C.ImU32(FieldColorFg)
	FieldTextOffset := self.FieldTextOffset

	result.TextOffset = C.int(FieldTextOffset)
	FieldClamp := self.FieldClamp

	result.Clamp = C.bool(FieldClamp)
	releaseFn = func() {
	}
	return result, releaseFn
}

func (self PlotAnnotation) c() (result C.ImPlotAnnotation, fin func()) {
	resultPtr, finFn := self.handle()
	return *resultPtr, finFn
}

func newPlotAnnotationFromC(cvalue *C.ImPlotAnnotation) *PlotAnnotation {
	result := new(PlotAnnotation)
	result.FieldPos = *(&Vec2{}).fromC(cvalue.Pos)
	result.FieldOffset = *(&Vec2{}).fromC(cvalue.Offset)
	result.FieldColorBg = uint32(cvalue.ColorBg)
	result.FieldColorFg = uint32(cvalue.ColorFg)
	result.FieldTextOffset = int32(cvalue.TextOffset)
	result.FieldClamp = cvalue.Clamp == C.bool(true)
	return result
}

type PlotAnnotationCollection struct {
	FieldAnnotations Vector[*PlotAnnotation]
	FieldTextBuffer  TextBuffer
	FieldSize        int32
}

func (self PlotAnnotationCollection) handle() (result *C.ImPlotAnnotationCollection, releaseFn func()) {
	result = new(C.ImPlotAnnotationCollection)
	FieldAnnotations := self.FieldAnnotations
	FieldAnnotationsData := FieldAnnotations.Data
	FieldAnnotationsDataArg, FieldAnnotationsDataFin := FieldAnnotationsData.handle()
	FieldAnnotationsVecArg := new(C.ImVector_ImPlotAnnotation)
	FieldAnnotationsVecArg.Size = C.int(FieldAnnotations.Size)
	FieldAnnotationsVecArg.Capacity = C.int(FieldAnnotations.Capacity)
	FieldAnnotationsVecArg.Data = FieldAnnotationsDataArg

	result.Annotations = *FieldAnnotationsVecArg
	FieldTextBuffer := self.FieldTextBuffer
	FieldTextBufferArg, FieldTextBufferFin := FieldTextBuffer.c()
	result.TextBuffer = FieldTextBufferArg
	FieldSize := self.FieldSize

	result.Size = C.int(FieldSize)
	releaseFn = func() {
		FieldAnnotationsDataFin()
		FieldTextBufferFin()
	}
	return result, releaseFn
}

func (self PlotAnnotationCollection) c() (result C.ImPlotAnnotationCollection, fin func()) {
	resultPtr, finFn := self.handle()
	return *resultPtr, finFn
}

func newPlotAnnotationCollectionFromC(cvalue *C.ImPlotAnnotationCollection) *PlotAnnotationCollection {
	result := new(PlotAnnotationCollection)
	result.FieldAnnotations = newVectorFromC(cvalue.Annotations.Size, cvalue.Annotations.Capacity, newPlotAnnotationFromC(cvalue.Annotations.Data))
	result.FieldTextBuffer = *newTextBufferFromC(&cvalue.TextBuffer)

	result.FieldSize = int32(cvalue.Size)
	return result
}

type PlotAxis struct {
	// TODO: contains unsupported fields
	data unsafe.Pointer
}

func (self PlotAxis) handle() (result *C.ImPlotAxis, releaseFn func()) {
	result = (*C.ImPlotAxis)(self.data)
	return result, func() {}
}

func (self PlotAxis) c() (result C.ImPlotAxis, fin func()) {
	resultPtr, finFn := self.handle()
	return *resultPtr, finFn
}

func newPlotAxisFromC(cvalue *C.ImPlotAxis) *PlotAxis {
	result := new(PlotAxis)
	result.data = unsafe.Pointer(cvalue)
	return result
}

type PlotColormapData struct {
	// TODO: contains unsupported fields
	data unsafe.Pointer
}

func (self PlotColormapData) handle() (result *C.ImPlotColormapData, releaseFn func()) {
	result = (*C.ImPlotColormapData)(self.data)
	return result, func() {}
}

func (self PlotColormapData) c() (result C.ImPlotColormapData, fin func()) {
	resultPtr, finFn := self.handle()
	return *resultPtr, finFn
}

func newPlotColormapDataFromC(cvalue *C.ImPlotColormapData) *PlotColormapData {
	result := new(PlotColormapData)
	result.data = unsafe.Pointer(cvalue)
	return result
}

type PlotContext struct {
	// TODO: contains unsupported fields
	data unsafe.Pointer
}

func (self PlotContext) handle() (result *C.ImPlotContext, releaseFn func()) {
	result = (*C.ImPlotContext)(self.data)
	return result, func() {}
}

func (self PlotContext) c() (result C.ImPlotContext, fin func()) {
	resultPtr, finFn := self.handle()
	return *resultPtr, finFn
}

func newPlotContextFromC(cvalue *C.ImPlotContext) *PlotContext {
	result := new(PlotContext)
	result.data = unsafe.Pointer(cvalue)
	return result
}

type PlotDateTimeSpec struct {
	// TODO: contains unsupported fields
	data unsafe.Pointer
}

func (self PlotDateTimeSpec) handle() (result *C.ImPlotDateTimeSpec, releaseFn func()) {
	result = (*C.ImPlotDateTimeSpec)(self.data)
	return result, func() {}
}

func (self PlotDateTimeSpec) c() (result C.ImPlotDateTimeSpec, fin func()) {
	resultPtr, finFn := self.handle()
	return *resultPtr, finFn
}

func newPlotDateTimeSpecFromC(cvalue *C.ImPlotDateTimeSpec) *PlotDateTimeSpec {
	result := new(PlotDateTimeSpec)
	result.data = unsafe.Pointer(cvalue)
	return result
}

type PlotInputMap struct {
	FieldPan           MouseButton
	FieldPanMod        int32
	FieldFit           MouseButton
	FieldSelect        MouseButton
	FieldSelectCancel  MouseButton
	FieldSelectMod     int32
	FieldSelectHorzMod int32
	FieldSelectVertMod int32
	FieldMenu          MouseButton
	FieldOverrideMod   int32
	FieldZoomMod       int32
	FieldZoomRate      float32
}

func (self PlotInputMap) handle() (result *C.ImPlotInputMap, releaseFn func()) {
	result = new(C.ImPlotInputMap)
	FieldPan := self.FieldPan

	result.Pan = C.ImGuiMouseButton(FieldPan)
	FieldPanMod := self.FieldPanMod

	result.PanMod = C.int(FieldPanMod)
	FieldFit := self.FieldFit

	result.Fit = C.ImGuiMouseButton(FieldFit)
	FieldSelect := self.FieldSelect

	result.Select = C.ImGuiMouseButton(FieldSelect)
	FieldSelectCancel := self.FieldSelectCancel

	result.SelectCancel = C.ImGuiMouseButton(FieldSelectCancel)
	FieldSelectMod := self.FieldSelectMod

	result.SelectMod = C.int(FieldSelectMod)
	FieldSelectHorzMod := self.FieldSelectHorzMod

	result.SelectHorzMod = C.int(FieldSelectHorzMod)
	FieldSelectVertMod := self.FieldSelectVertMod

	result.SelectVertMod = C.int(FieldSelectVertMod)
	FieldMenu := self.FieldMenu

	result.Menu = C.ImGuiMouseButton(FieldMenu)
	FieldOverrideMod := self.FieldOverrideMod

	result.OverrideMod = C.int(FieldOverrideMod)
	FieldZoomMod := self.FieldZoomMod

	result.ZoomMod = C.int(FieldZoomMod)
	FieldZoomRate := self.FieldZoomRate

	result.ZoomRate = C.float(FieldZoomRate)
	releaseFn = func() {
	}
	return result, releaseFn
}

func (self PlotInputMap) c() (result C.ImPlotInputMap, fin func()) {
	resultPtr, finFn := self.handle()
	return *resultPtr, finFn
}

func newPlotInputMapFromC(cvalue *C.ImPlotInputMap) *PlotInputMap {
	result := new(PlotInputMap)
	result.FieldPan = MouseButton(cvalue.Pan)
	result.FieldPanMod = int32(cvalue.PanMod)
	result.FieldFit = MouseButton(cvalue.Fit)
	result.FieldSelect = MouseButton(cvalue.Select)
	result.FieldSelectCancel = MouseButton(cvalue.SelectCancel)
	result.FieldSelectMod = int32(cvalue.SelectMod)
	result.FieldSelectHorzMod = int32(cvalue.SelectHorzMod)
	result.FieldSelectVertMod = int32(cvalue.SelectVertMod)
	result.FieldMenu = MouseButton(cvalue.Menu)
	result.FieldOverrideMod = int32(cvalue.OverrideMod)
	result.FieldZoomMod = int32(cvalue.ZoomMod)
	result.FieldZoomRate = float32(cvalue.ZoomRate)
	return result
}

type PlotItem struct {
	FieldID              ID
	FieldColor           uint32
	FieldLegendHoverRect Rect
	FieldNameOffset      int32
	FieldShow            bool
	FieldLegendHovered   bool
	FieldSeenThisFrame   bool
}

func (self PlotItem) handle() (result *C.ImPlotItem, releaseFn func()) {
	result = new(C.ImPlotItem)
	FieldID := self.FieldID

	result.ID = C.ImGuiID(FieldID)
	FieldColor := self.FieldColor

	result.Color = C.ImU32(FieldColor)
	FieldLegendHoverRect := self.FieldLegendHoverRect

	result.LegendHoverRect = FieldLegendHoverRect.toC()
	FieldNameOffset := self.FieldNameOffset

	result.NameOffset = C.int(FieldNameOffset)
	FieldShow := self.FieldShow

	result.Show = C.bool(FieldShow)
	FieldLegendHovered := self.FieldLegendHovered

	result.LegendHovered = C.bool(FieldLegendHovered)
	FieldSeenThisFrame := self.FieldSeenThisFrame

	result.SeenThisFrame = C.bool(FieldSeenThisFrame)
	releaseFn = func() {
	}
	return result, releaseFn
}

func (self PlotItem) c() (result C.ImPlotItem, fin func()) {
	resultPtr, finFn := self.handle()
	return *resultPtr, finFn
}

func newPlotItemFromC(cvalue *C.ImPlotItem) *PlotItem {
	result := new(PlotItem)
	result.FieldID = ID(cvalue.ID)
	result.FieldColor = uint32(cvalue.Color)
	result.FieldLegendHoverRect = *(&Rect{}).fromC(cvalue.LegendHoverRect)
	result.FieldNameOffset = int32(cvalue.NameOffset)
	result.FieldShow = cvalue.Show == C.bool(true)
	result.FieldLegendHovered = cvalue.LegendHovered == C.bool(true)
	result.FieldSeenThisFrame = cvalue.SeenThisFrame == C.bool(true)
	return result
}

type PlotItemGroup struct {
	// TODO: contains unsupported fields
	data unsafe.Pointer
}

func (self PlotItemGroup) handle() (result *C.ImPlotItemGroup, releaseFn func()) {
	result = (*C.ImPlotItemGroup)(self.data)
	return result, func() {}
}

func (self PlotItemGroup) c() (result C.ImPlotItemGroup, fin func()) {
	resultPtr, finFn := self.handle()
	return *resultPtr, finFn
}

func newPlotItemGroupFromC(cvalue *C.ImPlotItemGroup) *PlotItemGroup {
	result := new(PlotItemGroup)
	result.data = unsafe.Pointer(cvalue)
	return result
}

type PlotLegend struct {
	// TODO: contains unsupported fields
	data unsafe.Pointer
}

func (self PlotLegend) handle() (result *C.ImPlotLegend, releaseFn func()) {
	result = (*C.ImPlotLegend)(self.data)
	return result, func() {}
}

func (self PlotLegend) c() (result C.ImPlotLegend, fin func()) {
	resultPtr, finFn := self.handle()
	return *resultPtr, finFn
}

func newPlotLegendFromC(cvalue *C.ImPlotLegend) *PlotLegend {
	result := new(PlotLegend)
	result.data = unsafe.Pointer(cvalue)
	return result
}

type PlotNextItemData struct {
	// TODO: contains unsupported fields
	data unsafe.Pointer
}

func (self PlotNextItemData) handle() (result *C.ImPlotNextItemData, releaseFn func()) {
	result = (*C.ImPlotNextItemData)(self.data)
	return result, func() {}
}

func (self PlotNextItemData) c() (result C.ImPlotNextItemData, fin func()) {
	resultPtr, finFn := self.handle()
	return *resultPtr, finFn
}

func newPlotNextItemDataFromC(cvalue *C.ImPlotNextItemData) *PlotNextItemData {
	result := new(PlotNextItemData)
	result.data = unsafe.Pointer(cvalue)
	return result
}

type PlotNextPlotData struct {
	// TODO: contains unsupported fields
	data unsafe.Pointer
}

func (self PlotNextPlotData) handle() (result *C.ImPlotNextPlotData, releaseFn func()) {
	result = (*C.ImPlotNextPlotData)(self.data)
	return result, func() {}
}

func (self PlotNextPlotData) c() (result C.ImPlotNextPlotData, fin func()) {
	resultPtr, finFn := self.handle()
	return *resultPtr, finFn
}

func newPlotNextPlotDataFromC(cvalue *C.ImPlotNextPlotData) *PlotNextPlotData {
	result := new(PlotNextPlotData)
	result.data = unsafe.Pointer(cvalue)
	return result
}

type PlotPlot struct {
	// TODO: contains unsupported fields
	data unsafe.Pointer
}

func (self PlotPlot) handle() (result *C.ImPlotPlot, releaseFn func()) {
	result = (*C.ImPlotPlot)(self.data)
	return result, func() {}
}

func (self PlotPlot) c() (result C.ImPlotPlot, fin func()) {
	resultPtr, finFn := self.handle()
	return *resultPtr, finFn
}

func newPlotPlotFromC(cvalue *C.ImPlotPlot) *PlotPlot {
	result := new(PlotPlot)
	result.data = unsafe.Pointer(cvalue)
	return result
}

type PlotPointError struct {
	FieldX   float64
	FieldY   float64
	FieldNeg float64
	FieldPos float64
}

func (self PlotPointError) handle() (result *C.ImPlotPointError, releaseFn func()) {
	result = new(C.ImPlotPointError)
	FieldX := self.FieldX

	result.X = C.double(FieldX)
	FieldY := self.FieldY

	result.Y = C.double(FieldY)
	FieldNeg := self.FieldNeg

	result.Neg = C.double(FieldNeg)
	FieldPos := self.FieldPos

	result.Pos = C.double(FieldPos)
	releaseFn = func() {
	}
	return result, releaseFn
}

func (self PlotPointError) c() (result C.ImPlotPointError, fin func()) {
	resultPtr, finFn := self.handle()
	return *resultPtr, finFn
}

func newPlotPointErrorFromC(cvalue *C.ImPlotPointError) *PlotPointError {
	result := new(PlotPointError)
	result.FieldX = float64(cvalue.X)
	result.FieldY = float64(cvalue.Y)
	result.FieldNeg = float64(cvalue.Neg)
	result.FieldPos = float64(cvalue.Pos)
	return result
}

type PlotRange struct {
	FieldMin float64
	FieldMax float64
}

func (self PlotRange) handle() (result *C.ImPlotRange, releaseFn func()) {
	result = new(C.ImPlotRange)
	FieldMin := self.FieldMin

	result.Min = C.double(FieldMin)
	FieldMax := self.FieldMax

	result.Max = C.double(FieldMax)
	releaseFn = func() {
	}
	return result, releaseFn
}

func (self PlotRange) c() (result C.ImPlotRange, fin func()) {
	resultPtr, finFn := self.handle()
	return *resultPtr, finFn
}

func newPlotRangeFromC(cvalue *C.ImPlotRange) *PlotRange {
	result := new(PlotRange)
	result.FieldMin = float64(cvalue.Min)
	result.FieldMax = float64(cvalue.Max)
	return result
}

type PlotRect struct {
	FieldX PlotRange
	FieldY PlotRange
}

func (self PlotRect) handle() (result *C.ImPlotRect, releaseFn func()) {
	result = new(C.ImPlotRect)
	FieldX := self.FieldX
	FieldXArg, FieldXFin := FieldX.c()
	result.X = FieldXArg
	FieldY := self.FieldY
	FieldYArg, FieldYFin := FieldY.c()
	result.Y = FieldYArg
	releaseFn = func() {
		FieldXFin()
		FieldYFin()
	}
	return result, releaseFn
}

func (self PlotRect) c() (result C.ImPlotRect, fin func()) {
	resultPtr, finFn := self.handle()
	return *resultPtr, finFn
}

func newPlotRectFromC(cvalue *C.ImPlotRect) *PlotRect {
	result := new(PlotRect)
	result.FieldX = *newPlotRangeFromC(&cvalue.X)

	result.FieldY = *newPlotRangeFromC(&cvalue.Y)

	return result
}

type PlotStyle struct {
	// TODO: contains unsupported fields
	data unsafe.Pointer
}

func (self PlotStyle) handle() (result *C.ImPlotStyle, releaseFn func()) {
	result = (*C.ImPlotStyle)(self.data)
	return result, func() {}
}

func (self PlotStyle) c() (result C.ImPlotStyle, fin func()) {
	resultPtr, finFn := self.handle()
	return *resultPtr, finFn
}

func newPlotStyleFromC(cvalue *C.ImPlotStyle) *PlotStyle {
	result := new(PlotStyle)
	result.data = unsafe.Pointer(cvalue)
	return result
}

type PlotSubplot struct {
	// TODO: contains unsupported fields
	data unsafe.Pointer
}

func (self PlotSubplot) handle() (result *C.ImPlotSubplot, releaseFn func()) {
	result = (*C.ImPlotSubplot)(self.data)
	return result, func() {}
}

func (self PlotSubplot) c() (result C.ImPlotSubplot, fin func()) {
	resultPtr, finFn := self.handle()
	return *resultPtr, finFn
}

func newPlotSubplotFromC(cvalue *C.ImPlotSubplot) *PlotSubplot {
	result := new(PlotSubplot)
	result.data = unsafe.Pointer(cvalue)
	return result
}

type PlotTag struct {
	// TODO: contains unsupported fields
	data unsafe.Pointer
}

func (self PlotTag) handle() (result *C.ImPlotTag, releaseFn func()) {
	result = (*C.ImPlotTag)(self.data)
	return result, func() {}
}

func (self PlotTag) c() (result C.ImPlotTag, fin func()) {
	resultPtr, finFn := self.handle()
	return *resultPtr, finFn
}

func newPlotTagFromC(cvalue *C.ImPlotTag) *PlotTag {
	result := new(PlotTag)
	result.data = unsafe.Pointer(cvalue)
	return result
}

type PlotTagCollection struct {
	FieldTags       Vector[*PlotTag]
	FieldTextBuffer TextBuffer
	FieldSize       int32
}

func (self PlotTagCollection) handle() (result *C.ImPlotTagCollection, releaseFn func()) {
	result = new(C.ImPlotTagCollection)
	FieldTags := self.FieldTags
	FieldTagsData := FieldTags.Data
	FieldTagsDataArg, FieldTagsDataFin := FieldTagsData.handle()
	FieldTagsVecArg := new(C.ImVector_ImPlotTag)
	FieldTagsVecArg.Size = C.int(FieldTags.Size)
	FieldTagsVecArg.Capacity = C.int(FieldTags.Capacity)
	FieldTagsVecArg.Data = FieldTagsDataArg

	result.Tags = *FieldTagsVecArg
	FieldTextBuffer := self.FieldTextBuffer
	FieldTextBufferArg, FieldTextBufferFin := FieldTextBuffer.c()
	result.TextBuffer = FieldTextBufferArg
	FieldSize := self.FieldSize

	result.Size = C.int(FieldSize)
	releaseFn = func() {
		FieldTagsDataFin()
		FieldTextBufferFin()
	}
	return result, releaseFn
}

func (self PlotTagCollection) c() (result C.ImPlotTagCollection, fin func()) {
	resultPtr, finFn := self.handle()
	return *resultPtr, finFn
}

func newPlotTagCollectionFromC(cvalue *C.ImPlotTagCollection) *PlotTagCollection {
	result := new(PlotTagCollection)
	result.FieldTags = newVectorFromC(cvalue.Tags.Size, cvalue.Tags.Capacity, newPlotTagFromC(cvalue.Tags.Data))
	result.FieldTextBuffer = *newTextBufferFromC(&cvalue.TextBuffer)

	result.FieldSize = int32(cvalue.Size)
	return result
}

type PlotTick struct {
	FieldPlotPos    float64
	FieldPixelPos   float32
	FieldLabelSize  Vec2
	FieldTextOffset int32
	FieldMajor      bool
	FieldShowLabel  bool
	FieldLevel      int32
	FieldIdx        int32
}

func (self PlotTick) handle() (result *C.ImPlotTick, releaseFn func()) {
	result = new(C.ImPlotTick)
	FieldPlotPos := self.FieldPlotPos

	result.PlotPos = C.double(FieldPlotPos)
	FieldPixelPos := self.FieldPixelPos

	result.PixelPos = C.float(FieldPixelPos)
	FieldLabelSize := self.FieldLabelSize

	result.LabelSize = FieldLabelSize.toC()
	FieldTextOffset := self.FieldTextOffset

	result.TextOffset = C.int(FieldTextOffset)
	FieldMajor := self.FieldMajor

	result.Major = C.bool(FieldMajor)
	FieldShowLabel := self.FieldShowLabel

	result.ShowLabel = C.bool(FieldShowLabel)
	FieldLevel := self.FieldLevel

	result.Level = C.int(FieldLevel)
	FieldIdx := self.FieldIdx

	result.Idx = C.int(FieldIdx)
	releaseFn = func() {
	}
	return result, releaseFn
}

func (self PlotTick) c() (result C.ImPlotTick, fin func()) {
	resultPtr, finFn := self.handle()
	return *resultPtr, finFn
}

func newPlotTickFromC(cvalue *C.ImPlotTick) *PlotTick {
	result := new(PlotTick)
	result.FieldPlotPos = float64(cvalue.PlotPos)
	result.FieldPixelPos = float32(cvalue.PixelPos)
	result.FieldLabelSize = *(&Vec2{}).fromC(cvalue.LabelSize)
	result.FieldTextOffset = int32(cvalue.TextOffset)
	result.FieldMajor = cvalue.Major == C.bool(true)
	result.FieldShowLabel = cvalue.ShowLabel == C.bool(true)
	result.FieldLevel = int32(cvalue.Level)
	result.FieldIdx = int32(cvalue.Idx)
	return result
}

type PlotTicker struct {
	FieldTicks      Vector[*PlotTick]
	FieldTextBuffer TextBuffer
	FieldMaxSize    Vec2
	FieldLateSize   Vec2
	FieldLevels     int32
}

func (self PlotTicker) handle() (result *C.ImPlotTicker, releaseFn func()) {
	result = new(C.ImPlotTicker)
	FieldTicks := self.FieldTicks
	FieldTicksData := FieldTicks.Data
	FieldTicksDataArg, FieldTicksDataFin := FieldTicksData.handle()
	FieldTicksVecArg := new(C.ImVector_ImPlotTick)
	FieldTicksVecArg.Size = C.int(FieldTicks.Size)
	FieldTicksVecArg.Capacity = C.int(FieldTicks.Capacity)
	FieldTicksVecArg.Data = FieldTicksDataArg

	result.Ticks = *FieldTicksVecArg
	FieldTextBuffer := self.FieldTextBuffer
	FieldTextBufferArg, FieldTextBufferFin := FieldTextBuffer.c()
	result.TextBuffer = FieldTextBufferArg
	FieldMaxSize := self.FieldMaxSize

	result.MaxSize = FieldMaxSize.toC()
	FieldLateSize := self.FieldLateSize

	result.LateSize = FieldLateSize.toC()
	FieldLevels := self.FieldLevels

	result.Levels = C.int(FieldLevels)
	releaseFn = func() {
		FieldTicksDataFin()
		FieldTextBufferFin()
	}
	return result, releaseFn
}

func (self PlotTicker) c() (result C.ImPlotTicker, fin func()) {
	resultPtr, finFn := self.handle()
	return *resultPtr, finFn
}

func newPlotTickerFromC(cvalue *C.ImPlotTicker) *PlotTicker {
	result := new(PlotTicker)
	result.FieldTicks = newVectorFromC(cvalue.Ticks.Size, cvalue.Ticks.Capacity, newPlotTickFromC(cvalue.Ticks.Data))
	result.FieldTextBuffer = *newTextBufferFromC(&cvalue.TextBuffer)

	result.FieldMaxSize = *(&Vec2{}).fromC(cvalue.MaxSize)
	result.FieldLateSize = *(&Vec2{}).fromC(cvalue.LateSize)
	result.FieldLevels = int32(cvalue.Levels)
	return result
}
