package ray2sing

import (
	"testing"

	"github.com/sagernet/wireguard-go/donutvpn"
)

// Xray WARP links: wnoise=random is plain random bytes, wnoise=quic QUIC Initials. Both were mapped
// to a mode the WireGuard device does not know, so it sent random packets after a header.
func TestWireGuardXrayNoise(t *testing.T) {
	random := getWireGuardNoise(map[string]string{"wnoise": "random", "wnoisecount": "5", "wpayloadsize": "5-10", "wnoisedelay": "2-5"}, false)
	fake := random.FakePacket
	if !fake.Enabled || fake.Mode != donutvpn.NoiseModeRaw || fake.Count != (donutvpn.Range{From: 5, To: 5}) || fake.Size != (donutvpn.Range{From: 5, To: 10}) || fake.Delay != (donutvpn.Range{From: 2, To: 5}) {
		t.Fatalf("wnoise=random: %+v", fake)
	}
	quic := getWireGuardNoise(map[string]string{"wnoise": "quic", "wnoisecount": "2-4"}, false)
	if quic.FakePacket.Mode != donutvpn.NoiseModeQUIC {
		t.Fatalf("wnoise=quic: %+v", quic.FakePacket)
	}
	if mode := defaultWireguardNoiseOptions().FakePacket.Mode; mode != "" {
		t.Fatalf("default noise mode %q is not a mode", mode)
	}
}
