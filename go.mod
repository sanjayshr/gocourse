module github.com/sanjayshr/gocourse

require github.com/sanjayshr/gocourse/utils v0.0.0

require github.com/sanjayshr/gocourse/ninja v0.0.0

replace github.com/sanjayshr/gocourse/utils v0.0.0 => ./utils

replace github.com/sanjayshr/gocourse/ninja v0.0.0 => ./ninja

go 1.16
