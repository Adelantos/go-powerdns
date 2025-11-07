package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsReadOnlyHTTP(t *testing.T) {
	assert.True(t, IsReadOnlyHTTP(string(MDLuaAXFRScript)))
	assert.False(t, IsReadOnlyHTTP(string(MDAllowAXFRFrom)))
	assert.False(t, IsReadOnlyHTTP("CUSTOM"))
}

func TestIsNotViaHTTP(t *testing.T) {
	assert.True(t, IsNotViaHTTP(string(MDApiRectify)))
	assert.False(t, IsNotViaHTTP(string(MDLuaAXFRScript)))
	assert.False(t, IsNotViaHTTP("ALLOW-TSIG"))
}

func TestIsCustomKind(t *testing.T) {
	assert.True(t, IsCustomKind("X-App-Feature"))
	assert.True(t, IsCustomKind("x-other"))
	assert.False(t, IsCustomKind("App-Feature"))
	assert.False(t, IsCustomKind("XY"))
}
