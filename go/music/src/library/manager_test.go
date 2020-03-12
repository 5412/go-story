package library

import "testing"

func TestOps(t *testing.T) {
    mm := NewMusicManager()
    if mm == nil {
        t.Error("NewMusicManager is failed.")
    }
    if mm.Len() != 0 {
        t.Error("NewMusicManager failed, not empty")
    }

    m0 := &MusicEntry{
        "1",
        "My Heart Will Go On",
        "Celion Dion",
        "http://qbox.me/23501234",
        "MP3"}

    mm.Add(m0)

    if mm.Len() != 1 {
        t.Error("MusicManager Add() failed.")
    }

    m := mm.Find(m0.Name)
    if m == nil {
        t.Error("MusicManager.Find() failed.")
    }

    if m.Name != m0.Name || m.Id != m0.Id {
        t.Error("MusicManager.Find() failed. Found item mismatch.")
    }

    m, err := mm.Get(0)

    if m == nil {
        t.Error("MusicManager.Get() failed.", err)
    }

    m = mm.Remove(0)
    if m == nil || mm.Len() != 0 {
        t.Error("MusicManager.Remove() failed.",m)
    }
}

