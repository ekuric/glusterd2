package bricksplanner

import (
	smartvolapi "github.com/gluster/glusterd2/plugins/smartvol/api"
)

type disperseSubvolPlanner struct {
	subvolSize              uint64
	disperseCount           int
	disperseDataCount       int
	disperseRedundancyCount int
	brickSize               uint64
}

func (s *disperseSubvolPlanner) Init(req *smartvolapi.Volume, subvolSize uint64) {
	s.subvolSize = subvolSize
	s.disperseCount = req.DisperseCount
	s.disperseDataCount = req.DisperseDataCount
	s.disperseRedundancyCount = req.DisperseRedundancyCount
	s.brickSize = s.subvolSize / uint64(s.disperseDataCount)
}

func (s *disperseSubvolPlanner) BricksCount() int {
	return s.disperseCount
}

func (s *disperseSubvolPlanner) BrickSize(idx int) uint64 {
	return s.brickSize
}

func (s *disperseSubvolPlanner) BrickType(idx int) string {
	return "Brick"
}

func init() {
	subvolPlanners["disperse"] = &disperseSubvolPlanner{}
}
