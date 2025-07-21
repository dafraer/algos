type heap struct {
    h []int
}

func newHeap(s []int) *heap {
    sort.Sort(sort.Reverse(sort.IntSlice(s)))
    s = append([]int{0}, s...)
    return &heap{h:s}
}

func (h *heap) insert(num int) {
    h.h = append(h.h, num)
    i := len(h.h)-1
    for h.h[i] > h.h[i/2] && i/2 > 0 {
        h.h[i], h.h[i/2] = h.h[i/2], h.h[i]
        i/=2
    }
    fmt.Println("insert", h.h)
}

func (h *heap) pop() int {
    if len(h.h) <= 1 {
        return 0
    }

    top := h.h[1]
    h.h[1] = h.h[len(h.h)-1]  // Replace root with last element
    h.h = h.h[:len(h.h)-1]    // Shrink the slice

    // Heapify down
    i := 1
    for {
        left := 2 * i
        right := 2 * i + 1
        largest := i

        if left < len(h.h) && h.h[left] > h.h[largest] {
            largest = left
        }
        if right < len(h.h) && h.h[right] > h.h[largest] {
            largest = right
        }

        if largest == i {
            break
        }

        h.h[i], h.h[largest] = h.h[largest], h.h[i]
        i = largest
    }

    fmt.Println("pop", h.h, top)
    return top
}


func lastStoneWeight(stones []int) int {
    h := newHeap(stones)
    for len(h.h) > 2 {
        y := h.pop()
        x := h.pop()
        if y > x {
            h.insert(y-x)
        }
    }    
    if len(h.h) == 1 {
        return 0
    }
   return h.pop()
}
