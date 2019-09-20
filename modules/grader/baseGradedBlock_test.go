package grader_test

import (
	"testing"

	lxr "github.com/pegnet/LXRHash"

	. "github.com/pegnet/pegnet/modules/grader"
	"github.com/pegnet/pegnet/modules/testutils"
)

type OprData struct {
	EntryHash []byte
	ExtIDs    [][]byte
	Content   []byte
	Valid     bool
}

var x BlockGrader = nil

// TestBaseGradedBlock_Invalid tests various invalid sets and checks the resulting block
// has nil winners and is obviously a no win block
func TestBaseGradedBlock_Invalid(t *testing.T) {
	// TODO: Change the bitsize to something smaller. The grader needs to use the same one
	testutils.SetTestLXR(lxr.Init(lxr.Seed, 30, lxr.HashSize, lxr.Passes))

	t.Run("V1", func(t *testing.T) {
		testBaseGradedBlock_Invalid(t, 1)
	})
	t.Run("V2", func(t *testing.T) {
		testBaseGradedBlock_Invalid(t, 2)
	})
}

func testBaseGradedBlock_Invalid(t *testing.T, version uint8) {
	checkEmpty := func(g BlockGrader, reason string) {
		block := g.Grade()
		if block.Winners() != nil {
			t.Error(reason)
		}

		if block.Version() != version {
			t.Error("wrong version")
		}
	}

	dbht := int32(1)

	// Test the empty set
	t.Run("test the empty set", func(t *testing.T) {
		g, _ := NewGrader(version, dbht, nil)
		checkEmpty(g, "no oprs")
	})

	t.Run("test not enough winners", func(t *testing.T) {
		// Test an incomplete set
		winners := testutils.RandomWinners(version)
		g, _ := NewGrader(version, dbht, winners)
		for i := 0; i < testutils.WinnerAmt(version)-1; i++ {
			err := g.AddOPR(testutils.RandomOPRWithFields(version, dbht, winners))
			if err != nil {
				t.Error(err)
			}
		}

		if g.Count() != testutils.WinnerAmt(version)-1 {
			t.Errorf("exp %d count, found %d", testutils.WinnerAmt(version)-1, g.Count())
		}

		checkEmpty(g, "not enough oprs to have winners")
	})

	t.Run("test winner short hashes resorts to the previous", func(t *testing.T) {
		// Test an incomplete set
		winners := testutils.RandomWinners(version)
		g, _ := NewGrader(version, dbht, winners)
		block := g.Grade()
		if len(block.WinnersShortHashes()) != len(winners) {
			t.Error("short hashes did not grab the previous")
		} else {
			for i := range winners {
				if block.WinnersShortHashes()[i] != winners[i] {
					t.Error("wrong winner")
				}
			}
		}
	})

}

// TestBaseGradedBlock_Valid tests various valid sets and checks the resulting block
// has some winners and a good block. It will also check the cutoff and number of graded
// for sets of varying amounts.
func TestBaseGradedBlock_Valid(t *testing.T) {
	// TODO: Change the bitsize to something smaller. The grader needs to use the same one
	testutils.SetTestLXR(lxr.Init(lxr.Seed, 30, lxr.HashSize, lxr.Passes))

	t.Run("V1", func(t *testing.T) {
		testBaseGradedBlock_valid(t, 1)
	})
	t.Run("V2", func(t *testing.T) {
		testBaseGradedBlock_valid(t, 2)
	})
}

func testBaseGradedBlock_valid(t *testing.T, version uint8) {
	prevWinners := testutils.RandomWinners(version)
	dbht := int32(1)

	testAmt := func(amt int) {
		g, _ := NewGrader(version, dbht, prevWinners)
		for i := 0; i < amt; i++ {
			err := g.AddOPR(testutils.RandomOPRWithFields(version, dbht, prevWinners))
			if err != nil {
				t.Error(err)
			}
		}

		if g.Count() != amt {
			t.Errorf("exp %d count, found %d", amt, g.Count())
		}

		block := g.Grade()

		if block.Winners() == nil {
			t.Error("no winners found")
		}

		if len(block.Winners()) != testutils.WinnerAmt(version) {
			t.Error("not the right number of winners")
		} else {
			for i, sh := range block.WinnersShortHashes() {
				if sh == prevWinners[i] {
					t.Error("shorthashes showing previous winners instead of new winners")
				}
			}
		}

		if block.Version() != version {
			t.Error("wrong version")
		}

		cutoff := 50
		if amt < 50 {
			cutoff = amt
		}

		if block.Cutoff() != cutoff {
			t.Errorf("exp cutoff of %d, found %d", cutoff, block.Cutoff())
		}

		if len(block.Graded()) != cutoff {
			t.Errorf("exp graded of %d, found %d", cutoff, len(block.Graded()))
		}

		if amt > 50 {
			// Test a custom graded
			block = g.GradeCustom(amt)
			if block.Cutoff() != amt {
				t.Errorf("exp cutoff of %d, found %d", amt, block.Cutoff())
			}

			if len(block.Graded()) != amt {
				t.Errorf("exp graded of %d, found %d", amt, len(block.Graded()))
			}
		}

	}

	t.Run("test just enough winners", func(t *testing.T) {
		testAmt(testutils.WinnerAmt(version))
	})

	t.Run("test 35 oprs", func(t *testing.T) {
		testAmt(35)
	})

	t.Run("test 50 oprs", func(t *testing.T) {
		testAmt(50)
	})

	t.Run("test 100 oprs", func(t *testing.T) {
		testAmt(100)
	})

}
