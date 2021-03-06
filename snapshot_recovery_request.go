package raft

import (
	"code.google.com/p/goprotobuf/proto"
	"github.com/benbjohnson/go-raft/protobuf"
	"io"
	"io/ioutil"
)

// The request sent to a server to start from the snapshot.
type SnapshotRecoveryRequest struct {
	LeaderName string
	LastIndex  uint64
	LastTerm   uint64
	Peers      []string
	State      []byte
}

//------------------------------------------------------------------------------
//
// Constructors
//
//------------------------------------------------------------------------------

// Creates a new Snapshot request.
func newSnapshotRecoveryRequest(leaderName string, snapshot *Snapshot) *SnapshotRecoveryRequest {
	return &SnapshotRecoveryRequest{
		LeaderName: leaderName,
		LastIndex:  snapshot.LastIndex,
		LastTerm:   snapshot.LastTerm,
		Peers:      snapshot.Peers,
		State:      snapshot.State,
	}
}

// Encodes the SnapshotRecoveryRequest to a buffer. Returns the number of bytes
// written and any error that may have occurred.
func (req *SnapshotRecoveryRequest) encode(w io.Writer) (int, error) {

	p := proto.NewBuffer(nil)

	pb := &protobuf.ProtoSnapshotRecoveryRequest{
		LeaderName: proto.String(req.LeaderName),
		LastIndex:  proto.Uint64(req.LastIndex),
		LastTerm:   proto.Uint64(req.LastTerm),
		Peers:      req.Peers,
		State:      req.State,
	}
	err := p.Marshal(pb)

	if err != nil {
		return -1, err
	}

	return w.Write(p.Bytes())
}

// Decodes the SnapshotRecoveryRequest from a buffer. Returns the number of bytes read and
// any error that occurs.
func (req *SnapshotRecoveryRequest) decode(r io.Reader) (int, error) {
	data, err := ioutil.ReadAll(r)

	if err != nil {
		return 0, err
	}

	totalBytes := len(data)

	pb := &protobuf.ProtoSnapshotRequest{}
	p := proto.NewBuffer(data)

	err = p.Unmarshal(pb)
	if err != nil {
		return -1, err
	}

	req.LeaderName = pb.GetLeaderName()
	req.LastIndex = pb.GetLastIndex()
	req.LastTerm = pb.GetLastTerm()
	req.Peers = req.Peers
	req.State = req.State

	return totalBytes, nil
}
