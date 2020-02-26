package kv

import (
	"context"
	"flag"
	"reflect"
	"testing"
	"time"

	"github.com/prysmaticlabs/prysm/slasher/flags"
	"github.com/urfave/cli"
)

type spanMapTestStruct struct {
	epoch   uint64
	spanMap map[uint64][2]uint16
}

var spanTests []spanMapTestStruct

func init() {
	spanTests = []spanMapTestStruct{
		{
			epoch: 1,
			spanMap: map[uint64][2]uint16{
				1: {10, 20},
				2: {11, 21},
				3: {12, 22},
			},
		},
		{
			epoch: 2,
			spanMap: map[uint64][2]uint16{
				1: {10, 20},
				2: {11, 21},
				3: {12, 22},
			},
		},
		{
			epoch: 3,
			spanMap: map[uint64][2]uint16{
				1: {10, 20},
				2: {11, 21},
				3: {12, 22},
			},
		},
	}
}

func TestValidatorSpanMap_NilDB(t *testing.T) {
	app := cli.NewApp()
	set := flag.NewFlagSet("test", 0)
	db := setupDB(t, cli.NewContext(app, set, nil))
	defer teardownDB(t, db)
	ctx := context.Background()

	validatorIdx := uint64(1)
	vsm, err := db.EpochSpansMap(ctx, validatorIdx)
	if err != nil {
		t.Fatalf("Nil EpochSpansMap should not return error: %v", err)
	}
	if !reflect.DeepEqual(vsm, map[uint64][2]uint16{}) {
		t.Fatal("EpochSpansMap should return nil")
	}
}

func TestValidatorSpanMap_Save(t *testing.T) {
	app := cli.NewApp()
	set := flag.NewFlagSet("test", 0)
	db := setupDB(t, cli.NewContext(app, set, nil))
	defer teardownDB(t, db)
	ctx := context.Background()

	for _, tt := range spanTests {
		err := db.SaveEpochSpansMap(ctx, tt.epoch, tt.spanMap)
		if err != nil {
			t.Fatalf("Save validator span map failed: %v", err)
		}
		sm, err := db.EpochSpansMap(ctx, tt.epoch)
		if err != nil {
			t.Fatalf("Failed to get validator span map: %v", err)
		}

		if sm == nil || !reflect.DeepEqual(sm, tt.spanMap) {
			t.Fatalf("Get should return validator span map: %v got: %v", tt.spanMap, sm)
		}
		s, err := db.EpochSpanByValidatorIndex(ctx, tt.epoch, 1)
		if err != nil {
			t.Fatalf("Failed to get validator span for epoch 1: %v", err)
		}
		if s == [2]uint16{} || !reflect.DeepEqual(s, tt.spanMap[1]) {
			t.Fatalf("Get should return validator spans for epoch 1: %v got: %v", tt.spanMap[1], s)
		}
	}
}

func TestValidatorSpanMap_SaveWithCache(t *testing.T) {
	app := cli.NewApp()
	set := flag.NewFlagSet("test", 0)
	set.Bool(flags.UseSpanCacheFlag.Name, true, "enable span map cache")
	db := setupDB(t, cli.NewContext(app, set, nil))
	defer teardownDB(t, db)
	ctx := context.Background()

	for _, tt := range spanTests {
		err := db.SaveEpochSpansMap(ctx, tt.epoch, tt.spanMap)
		if err != nil {
			t.Fatalf("Save validator span map failed: %v", err)
		}
		// wait for value to pass through cache buffers
		time.Sleep(time.Millisecond * 10)
		sm, err := db.EpochSpansMap(ctx, tt.epoch)
		if err != nil {
			t.Fatalf("Failed to get validator span map: %v", err)
		}

		if sm == nil || !reflect.DeepEqual(sm, tt.spanMap) {
			t.Fatalf("Get should return validator span map: %v got: %v", tt.spanMap, sm)
		}
	}
}

func TestValidatorSpanMap_Delete(t *testing.T) {
	app := cli.NewApp()
	set := flag.NewFlagSet("test", 0)
	db := setupDB(t, cli.NewContext(app, set, nil))
	defer teardownDB(t, db)
	ctx := context.Background()

	for _, tt := range spanTests {
		err := db.SaveEpochSpansMap(ctx, tt.epoch, tt.spanMap)
		if err != nil {
			t.Fatalf("Save validator span map failed: %v", err)
		}
	}

	for _, tt := range spanTests {
		sm, err := db.EpochSpansMap(ctx, tt.epoch)
		if err != nil {
			t.Fatalf("Failed to get validator span map: %v", err)
		}
		if sm == nil || !reflect.DeepEqual(sm, tt.spanMap) {
			t.Fatalf("Get should return validator span map: %v got: %v", tt.spanMap, sm)
		}
		err = db.DeleteEpochSpans(ctx, tt.epoch)
		if err != nil {
			t.Fatalf("Delete validator span map error: %v", err)
		}
		sm, err = db.EpochSpansMap(ctx, tt.epoch)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(sm, map[uint64][2]uint16{}) {
			t.Errorf("Expected validator span map to be deleted, received: %v", sm)
		}
	}
}

func TestValidatorSpanMap_DeleteWithCache(t *testing.T) {
	app := cli.NewApp()
	set := flag.NewFlagSet("test", 0)
	set.Bool(flags.UseSpanCacheFlag.Name, true, "enable span map cache")
	db := setupDB(t, cli.NewContext(app, set, nil))
	defer teardownDB(t, db)
	ctx := context.Background()

	for _, tt := range spanTests {
		err := db.SaveEpochSpansMap(ctx, tt.epoch, tt.spanMap)
		if err != nil {
			t.Fatalf("Save validator span map failed: %v", err)
		}
	}
	// wait for value to pass through cache buffers
	time.Sleep(time.Millisecond * 10)
	for _, tt := range spanTests {
		sm, err := db.EpochSpansMap(ctx, tt.epoch)
		if err != nil {
			t.Fatalf("Failed to get validator span map: %v", err)
		}
		if sm == nil || !reflect.DeepEqual(sm, tt.spanMap) {
			t.Fatalf("Get should return validator span map: %v got: %v", tt.spanMap, sm)
		}
		err = db.DeleteEpochSpans(ctx, tt.epoch)
		if err != nil {
			t.Fatalf("Delete validator span map error: %v", err)
		}
		// wait for value to pass through cache buffers
		time.Sleep(time.Millisecond * 10)
		sm, err = db.EpochSpansMap(ctx, tt.epoch)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(sm, map[uint64][2]uint16{}) {
			t.Errorf("Expected validator span map to be deleted, received: %v", sm)
		}
	}
}

//func TestValidatorSpanMap_SaveOnEvict(t *testing.T) {
//	db := setupDBDiffCacheSize(t, 5, 5)
//	defer teardownDB(t, db)
//	ctx := context.Background()
//
//	tsm := &spanMapTestStruct{
//		epoch: 1,
//		spanMap: map[uint64][2]uint16{
//			1: {10, 20},
//			2: {11, 21},
//			3: {12, 22},
//		},
//	}
//	for i := uint64(0); i < 6; i++ {
//		err := db.SaveEpochSpansMap(ctx, i, tsm.spanMap)
//		if err != nil {
//			t.Fatalf("Save validator span map failed: %v", err)
//		}
//	}
//
//	// Wait for value to pass through cache buffers.
//	time.Sleep(time.Millisecond * 1000)
//	for i := uint64(0); i < 6; i++ {
//		sm, err := db.EpochSpansMap(ctx, i)
//		if err != nil {
//			t.Fatalf("Failed to get validator span map: %v", err)
//		}
//		if sm == nil || !reflect.DeepEqual(sm, tsm.spanMap) {
//			t.Fatalf("Get should return validator: %d span map: %v got: %v", i, tsm.spanMap, sm)
//		}
//	}
//}

//func TestValidatorSpanMap_SaveCachedSpansMaps(t *testing.T) {
//	app := cli.NewApp()
//	set := flag.NewFlagSet("test", 0)
//	set.Bool(flags.UseSpanCacheFlag.Name, true, "enable span map cache")
//	db := setupDB(t, cli.NewContext(app, set, nil))
//	defer teardownDB(t, db)
//	ctx := context.Background()
//
//	for _, tt := range spanTests {
//		err := db.SaveEpochSpansMap(ctx, tt.epoch, tt.spanMap)
//		if err != nil {
//			t.Fatalf("Save validator span map failed: %v", err)
//		}
//	}
//	// wait for value to pass through cache buffers
//	time.Sleep(time.Millisecond * 10)
//	err := db.SaveCachedSpansMaps(ctx)
//	if err != nil {
//		t.Errorf("Failed to save cached span maps to db: %v", err)
//	}
//	db.spanCache.Clear()
//	for _, tt := range spanTests {
//		sm, err := db.EpochSpansMap(ctx, tt.epoch)
//		if err != nil {
//			t.Fatalf("Failed to get validator span map: %v", err)
//		}
//		if sm == nil || !proto.Equal(sm, tt.spanMap) {
//			t.Fatalf("Get should return validator span map: %v got: %v", tt.spanMap, sm)
//		}
//	}
//}
