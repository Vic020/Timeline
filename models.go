package main

import (
	"errors"
	"reflect"
)

type ItemCon struct {
	IndexNum int
	Items    []Item
}

func (ic *ItemCon) New(name, avatar, timestamp, content string) (Item, error) {
	index := ic.IndexNum + 1
	newItem := Item{
		index:     index,
		Name:      name,
		Avatar:    avatar,
		Timestamp: timestamp,
		Content:   content,
	}
	ic.Items = append(ic.Items, newItem)
	ic.IndexNum = index

	return newItem, nil
}

func (ic *ItemCon) Update(index int, name, avatar, timestamp, content string) (Item, error) {
	if index > ic.IndexNum {
		return Item{}, errors.New(ObjectNotExistError)
	}
	for _, item := range ic.Items {
		if item.index == index {
			item.Name = name
			item.Avatar = avatar
			item.Timestamp = timestamp
			item.Content = content
			return item, nil
		}
	}

	return Item{}, errors.New(IndexError)
}

func (ic *ItemCon) Delete(index int) error {
	for _, item := range ic.Items {
		if item.index == index {
			//ic.Items
			return nil
		}
	}
	return errors.New(IndexError)
}

func (ic *ItemCon) Get(index int) (Item, error) {
	for _, item := range ic.Items {
		if item.index == index {
			return item, nil
		}
	}
	return Item{}, errors.New(IndexError)
}

func (ic *ItemCon) All() []Item {
	if ic.Items == nil {
		return []Item{}
	} else {
		return ic.Items
	}
}

func (ic *ItemCon) List(start, end int) []Item {
	if start > len(ic.Items) {
		start = len(ic.Items)
	}
	if end > len(ic.Items) {
		end = len(ic.Items)
	}
	if ic.Items == nil {
		return []Item{}
	} else {
		return ic.Items[start:end]
	}
}

type Item struct {
	index     int    `json:"index"`
	Name      string `json:"name"`
	Avatar    string `json:"avatar"`
	Timestamp string `json:"timestamp"`
	Content   string `json:"content"`
}

func (item Item) IsEmpty() bool {
	return reflect.DeepEqual(item, Item{})
}
