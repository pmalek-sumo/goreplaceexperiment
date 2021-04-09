module replace1

go 1.16

require go.opentelemetry.io/collector/processor/filterprocessor v0.0.0

replace go.opentelemetry.io/collector/processor/filterprocessor => ./collector/processor/filterprocessor

// replace go.opentelemetry.io/collector => ./collector
