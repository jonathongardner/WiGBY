import argparse
import asyncio
import json
import logging
import os
import platform
import ssl

from aiohttp import web

from aiortc import RTCPeerConnection, RTCSessionDescription
from aiortc.exceptions import InvalidStateError
from aiortc.contrib.media import MediaPlayer, MediaRelay

ROOT = os.path.join(os.path.dirname(__file__), 'public')


relay = None
webcam = None


def create_local_tracks(play_from):
    global relay, webcam

    if play_from:
        player = MediaPlayer(play_from)
        return player.video
    else:
        options = {"framerate": "30", "video_size": "640x480"}
        if relay is None:
            webcam = MediaPlayer("/dev/video0", format="v4l2", options=options)
            relay = MediaRelay()
        return relay.subscribe(webcam.video)


async def index(request):
    content = open(os.path.join(ROOT, "index.html"), "r").read()
    return web.Response(content_type="text/html", text=content)


async def javascript(request):
    content = open(os.path.join(ROOT, "client.js"), "r").read()
    return web.Response(content_type="application/javascript", text=content)


async def offer(request):
    params = await request.json()
    offer = RTCSessionDescription(sdp=params["sdp"], type=params["type"])

    pc = RTCPeerConnection()
    pcHash = { 'channels': {} }
    pcs[pc] = pcHash

    @pc.on("connectionstatechange")
    async def on_connectionstatechange():
        print("Connection state is %s" % pc.connectionState)
        if pc.connectionState == "failed":
            await pc.close()
            if not pcs.pop(pc, None):
                print("No pc")

    @pc.on("datachannel")
    def on_datachannel(channel):
        print("Channel %s" % channel.label)
        if channel.label == 'car':
            pcHash['channels'][channel.label] = channel

    # open media source
    video = create_local_tracks(args.play_from)

    await pc.setRemoteDescription(offer)
    for t in pc.getTransceivers():
        if t.kind == "video" and video:
            pc.addTrack(video)

    answer = await pc.createAnswer()
    await pc.setLocalDescription(answer)

    return web.Response(
        content_type="application/json",
        text=json.dumps(
            {"sdp": pc.localDescription.sdp, "type": pc.localDescription.type}
        ),
    )


async def send_message(request):
    for (pc, pcHash) in pcs.items():
        if 'car' in pcHash['channels']:
            try:
                channel = pcHash['channels']['car']
                channel.send(
                    json.dumps({ "now": "cool" })
                )
                await channel._RTCDataChannel__transport._data_channel_flush()
                await channel._RTCDataChannel__transport._transmit()
            except InvalidStateError:
                print('Connection stale')
    return web.Response(
        content_type="application/json",
        text=json.dumps({ 'count': len(pcs) }),
    )


async def pcs_connections(request):
    return web.Response(
        content_type="application/json",
        text=json.dumps({ 'count': len(pcs) }),
    )


pcs = {}

async def on_shutdown(app):
    # close peer connections
    coros = [pc.close() for pc in pcs]
    await asyncio.gather(*coros)
    pcs.clear()


if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="WebRTC webcam demo")
    parser.add_argument("--play-from", help="Read the media from a file and sent it."),
    parser.add_argument(
        "--host", default="0.0.0.0", help="Host for HTTP server (default: 0.0.0.0)"
    )
    parser.add_argument(
        "--port", type=int, default=3000, help="Port for HTTP server (default: 8080)"
    )
    parser.add_argument("--verbose", "-v", action="count")
    args = parser.parse_args()

    if args.verbose:
        logging.basicConfig(level=logging.DEBUG)
    else:
        logging.basicConfig(level=logging.INFO)

    app = web.Application()
    app.on_shutdown.append(on_shutdown)
    app.router.add_get("/", index)
    app.router.add_get("/client.js", javascript)
    app.router.add_get("/api/pcs", pcs_connections)
    app.router.add_get("/api/send_message", send_message)
    app.router.add_post("/api/offer", offer)
    web.run_app(app, host=args.host, port=args.port)
