package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFizz(t *testing.T) {
	// Given
	input := 3

	// When
	result := fizzbuzz(input)

	// Then
	require.Equal(t, "fizz", result)
}

func TestBuzz(t *testing.T) {
	// GIVEN
	var input int
	input = 5

	// WHEN
	result := fizzbuzz(input)

	// THEN
	require.Equal(t, "buzz", result)
}

func TestFizzBuzz(t *testing.T) {
	// GIVEN
	input := 15

	// WHEN
	result := fizzbuzz(input)

	// THEN
	require.Equal(t, "fizzbuzz", result)
}

func TestDefault(t *testing.T) {
	// GIVEN
	input := 2

	// WHEN
	result := fizzbuzz(input)

	// THEN
	require.Equal(t, "2", result)
}

