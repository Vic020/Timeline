package main

import (
	"errors"
	"reflect"
)

type ItemCon struct {
	indexNum int
	items    []Item
}

func (ic *ItemCon) New(name, avatar, timestamp, content string) (Item, error) {
	index := ic.indexNum + 1
	newItem := Item{
		index:     index,
		Name:      name,
		Avatar:    avatar,
		Timestamp: timestamp,
		Content:   content,
	}
	ic.items = append(ic.items, newItem)
	ic.indexNum = index

	return newItem, nil
}

func (ic *ItemCon) Update(index int, name, avatar, timestamp, content string) (Item, error) {
	if index > ic.indexNum {
		return Item{}, errors.New(ObjectNotExistError)
	}
	for _, item := range ic.items {
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
	for _, item := range ic.items {
		if item.index == index {
			//ic.items
			return nil
		}
	}
	return errors.New(IndexError)
}

func (ic *ItemCon) Get(index int) (Item, error) {
	for _, item := range ic.items {
		if item.index == index {
			return item, nil
		}
	}
	return Item{}, errors.New(IndexError)
}

func (ic *ItemCon) All() []Item {
	if ic.items == nil {
		return []Item{}
	} else {
		return ic.items
	}
}

func (ic *ItemCon) List(offset, limit int) []Item {
	if ic.items == nil {
		return []Item{}
	} else {
		return ic.items[offset : offset+limit]
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
