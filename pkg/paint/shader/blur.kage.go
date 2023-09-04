//go:build ignore

package main

// BLUR RADIUS
var Radius float = 40.0

const Directions = 32
const Quality = 3
const Pi = 6.28318530718

func Fragment(position vec4, texCoord vec2, color vec4) vec4 {
	const deltaD = Pi / Directions    // [0,360]
	const deltaFactor = 1.0 / Quality // [0,1]

	radius := Radius

	uv := position.xy

	clr := imageColorAtPixel(uv)

	for d := 0.0; d < Pi; d += deltaD {
		for i := deltaFactor; i <= 1.001; i += deltaFactor {
			clr += imageColorAtPixel(uv + vec2(cos(d), sin(d))*radius*i)
		}
	}

	clr /= Quality*Directions - 15
	return clr
}

func imageColorAtPixel(pixelCoords vec2) vec4 {
	sizeInPixels := imageSrcTextureSize()
	originInTexels, _ := imageSrcRegionOnTexture()
	adjustedTexelCoords := pixelCoords/sizeInPixels + originInTexels
	return imageSrc0At(adjustedTexelCoords)
}
