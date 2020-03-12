package library

import "errors"

type MusicEntry struct {
    Id string
    Name string
    Artist string
    Source string
    Type string
}

type MusicManager struct {
    musics []MusicEntry
}

func NewMusicManager() *MusicManager {
    return &MusicManager{make([]MusicEntry, 0)}
}

func (m *MusicManager) Len() int {
    return len(m.musics)
}

func (m *MusicManager) Get(index int) (music *MusicEntry, err error) {
    if index < 0 || index >= m.Len() {
        return nil, errors.New("Index out of range.")
    }
    return &m.musics[index], nil
}

func (m *MusicManager) Find(name string) *MusicEntry {
    if m.Len() == 0 {
        return nil
    }
    // 注意 := range
    for _, music := range m.musics {
        if music.Name == name {
            return &music
        }
    }

    return nil
}

func (m *MusicManager) Add(music *MusicEntry) {
    m.musics = append(m.musics, *music)
}

func (m *MusicManager) Remove(index int) *MusicEntry {
    if index < 0 || index >= m.Len() {
        return nil
    }
    removeMusic := m.musics[index]
    if index < m.Len() - 1 { // middle element
        /***************************************************************/
        /** append 函数详解：
        /** 函数作用是为一个切片追加元素，如果该切片存储空间（cap）足够
        /** 直接追加，长度（len）变长，如果空间不足，就会重新开辟内存，并
        /** 将之前的元素和新的元素一同拷贝进去，第一个元素为切片自身
        /** 之后参数可以是多个与切片同类型的变量，也可以是另外一个切片，
        /** 如果是另外一个切片append 函数参数只能接收两个切片，并且末尾要
        /** 加三个点（.）
        /** 特殊用法，将字符串作为第二个参数传入append([]byte("hello"), "world"...)
        /***************************************************************/
        m.musics = append(m.musics[:index-1], m.musics[index + 1 :]... )
    } else if index == 0 { // frist element
        m.musics = make([]MusicEntry, 0)
    } else { // last element
        m.musics = m.musics[:index-1]
    }

    // 注意返回的参数类型
    return &removeMusic
}

