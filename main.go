package main

import (
	"context"
	"sync"
	"time"

	"github.com/Vardhanb07/CNP/base"
	"github.com/Vardhanb07/CNP/node"
)

func main() {
	end1node := node.NewEndNode(1, "end-1")
	end2node := node.NewEndNode(2, "end-2")
	end3node := node.NewEndNode(3, "end-3")
	router1node := node.NewRouterNode("router-1", 500)
	end1node.SetNext(router1node)
	end2node.SetNext(router1node)
	end3node.SetNext(router1node)
	end4node := node.NewEndNode(4, "end-4")
	router1node.SetNext(end4node)
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*3))
	defer cancel()
	wg.Add(4)
	// end1node
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				p := base.NewPacket([]byte("transfer from end1node to router1node"))
				end1node.Transfer(p)
			}
		}
	}()
	// end2node
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				p := base.NewPacket([]byte("transfer from end2node to router1node"))
				end2node.Transfer(p)
			}
		}
	}()
	//end3node
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				p := base.NewPacket([]byte("transfer from end3node to router1node"))
				end3node.Transfer(p)
			}
		}
	}()
	// router1node
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				p := base.NewPacket([]byte("transfer from router1node to end4node"))
				router1node.Transfer(p)
			}
		}
	}()
	wg.Wait()
}
