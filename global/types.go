package global

// Filters - фильтры для запросов
type Filters map[string]string

// Sorts - сортировка для запросов
type Sorts []string

// OptionsList - лист опций для выборок
type OptionsList struct {
	Offset  uint64
	Limit   uint64
	Filters Filters
	Sorts   Sorts
}
