/*
 * This content is Auto-Generated by cimgui-go's Make file.
 * DO NOT EDIT
 *
 * cimgui-go: https://github.com/gabstv/cimgui-go
 * From: templates/assert.h
 *
 * TODO:
 * - tbh we can return data about an error back to GO and let it panic from GO.
 * - figure out how to convert cond into an human-readable expression (like in default assert func)
 */
#define IM_ASSERT(_EXPR)  GoFriendlyAssert(_EXPR, __FILE__, __LINE__)
#include <stdio.h>
#include <stdlib.h>
static void GoFriendlyAssert(bool cond/*const char *expr*/, const char *file, int line) {
        if (cond) {
                return;
        }

        fprintf(stdout, "File: %s, Line: %d\n", file, line);

        // Instead of throwing an exception, we're using exit(1) to terminate the program.
        exit(1);
}

