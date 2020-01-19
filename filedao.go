package main

import "log"

type FileDao struct {
	fileSaver *FileSaver
	itemCon   *ItemCon
}

var FileDaoHandle FileDao

func initFileDao() error {
	FileDaoHandle = FileDao{fileSaver: &FileSaver{filePath: "timeline.db"}}

	decoder, err := FileDaoHandle.fileSaver.Load()
	if err != nil {
		if err.Error() == FileNotExistError {
			log.Println("Init DB")

			FileDaoHandle.itemCon = &ItemCon{items: nil, indexNum: 0}
			return nil
		}
		return err
	} else {
		if err := decoder.Decode(&FileDaoHandle.itemCon); err != nil {
			return err
		}

	}

	return nil
}

func (f *FileDao) GetItemList(page, limit int) []Item {
	return f.itemCon.List(page, limit)
}

func (f *FileDao) AddItem(name, avatar, timestamp, content string) Item {
	item, err := f.itemCon.New(name, avatar, timestamp, content)
	if err != nil {
		f.fileSaver.Store(f.itemCon)
		return Item{}
	} else {
		return item
	}
}

func (f *FileDao) UpdateItem(index int, name, avator, timestamp, content string) Item {
	item, err := f.itemCon.Update(index, name, avator, timestamp, content)

	if err != nil {
		f.fileSaver.Store(f.itemCon)
		return Item{}
	}

	return item
}
