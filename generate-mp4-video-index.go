package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pillash/mp4util"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
//--^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
// press l key for looping the video indefinitely
//--^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
//---------------------------------------------------------
//--^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
// use global Shortcuts
//--^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
//---------------------------------------------------------
//--^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
// + use localStorage to remember where you have been last
// - make it load exactly where you left it
//--^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
// + calculate the total amount of time in all files in index.html file
// - show how much time i have watched and how much is remaining till the end
// + also key n for next video and key p for previous video
// + left arrawy seek 5 sec to the left right arrow seek 5 to the right
// + press f to toggle fullscreen mode
// + space - pause video
// + pageup and pagedown for speed up/down the video playback
// + convert from seconds (player.duration) to minutes and seconds
// + Math.floor(Math.floor(player.duration)/60) + ":" + Math.floor(player.duration)%60

// player.addEventListener('durationchange', function() {
//     console.log('Duration change', player.duration);
// });

//--^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
// install the program
// go install github.com\lanastasov\generate-mp4-video-index
//--^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
//

func main() {

	fmt.Println("--generate-mp4-video-index.go-v-0.1.1")
	// https://github.com/lanastasov/add-zero-to-chapters-subfolder
	
	files, _ := filepath.Glob("*.mp4")
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	chapter := strings.Split(dir, "\\")

	if len(files) > 0 {
		f, err := os.Create("index.html")
		check(err)
		defer f.Close()

		_, err = f.WriteString("<!DOCTYPE html>\n")
		// check(err)
		_, err = f.WriteString("<html lang=\"en\">\n")
		_, err = f.WriteString("<head>\n")
		_, err = f.WriteString(fmt.Sprintf("<title>%s</title>\n", chapter[len(chapter)-1]))
		_, err = f.WriteString("<body>\n")
		_, err = f.WriteString("<video id=\"video\"  width=\"320\" height=\"240\" controls autoplay>\n")
		_, err = f.WriteString("<source src=\"./")
		_, err = f.WriteString(files[0])
		_, err = f.WriteString("\" type=\"video/mp4\">\n")
		_, err = f.WriteString("</video>\n")
		_, err = f.WriteString("<script>\n")
		_, err = f.WriteString("window.generate_mp4_video_index = 1;\n")
		_, err = f.WriteString("var counter = 0;\n")
		_, err = f.WriteString("var arr = [\n")
		totalFolderTime := 0
		for i := range files {
			val, _ := mp4util.Duration(files[i])
			totalFolderTime += val
			fmt.Println(files[i], "---", val);
			_, err = f.WriteString("\"./" + strings.ReplaceAll(files[i], "#", "%23") + "\",\n")
		}

		var h = int(totalFolderTime / 3600)
		var m = (totalFolderTime - int(totalFolderTime/3600)*3600) / 60
		var s = totalFolderTime - h*3600 - m*60

		_, err = f.WriteString("]\n")
		_, err = f.WriteString("var player=document.getElementsByTagName('video')[0];\n")
		_, err = f.WriteString("player.addEventListener('ended',myHandler,false);\n")
		_, err = f.WriteString("var fullScrn = true;\n")
		_, err = f.WriteString("const isVideoPlaying = video => !!(video.currentTime > 0 && !video.paused && !video.ended && video.readyState > 2);\n")
		_, err = f.WriteString("var playbackRate = 1.75;\n")
		_, err = f.WriteString("player.playbackRate = playbackRate;\n")
		_, err = f.WriteString("function togglePlay() {\n")
		_, err = f.WriteString("    if (isVideoPlaying(player)) {\n")
		_, err = f.WriteString("    setTimeout(function() {\n")
		_, err = f.WriteString("        player.pause();\n")
		_, err = f.WriteString("    },200);\n")
		_, err = f.WriteString("    } else {\n")
		_, err = f.WriteString("     setTimeout(function() {\n")
		_, err = f.WriteString("        player.play();\n")
		_, err = f.WriteString("     },200);\n")
		_, err = f.WriteString("    }\n")
		_, err = f.WriteString("}\n")
		_, err = f.WriteString("function myHandler(e) {\n")
		
		_, err = f.WriteString("if (e === '+') {\n")
		_, err = f.WriteString("	counter++;\n")
		_, err = f.WriteString("} else if (e === '-') {\n")
		_, err = f.WriteString("	counter--;\n")
		_, err = f.WriteString("} else {\n")
		_, err = f.WriteString("	counter++;\n")
		_, err = f.WriteString("}\n")
		_, err = f.WriteString("var videoSrcTitle = 'videoSrc_'+document.title;\n")
		_, err = f.WriteString("var videoSrc = localStorage.getItem(videoSrcTitle);\n")
		_, err = f.WriteString("if (counter >= 0 && counter <= arr.length-1) {\n")
		_, err = f.WriteString("	if (videoSrc === null || videoSrc === 'undefined') {\n")
		_, err = f.WriteString("		player.src = arr[counter];\n")
		_, err = f.WriteString("		localStorage.setItem(videoSrcTitle, arr[counter]);\n")
		_, err = f.WriteString("	} else {\n")
		_, err = f.WriteString("		if (counter !== arr.indexOf(videoSrc)+1) {\n")
		_, err = f.WriteString("			counter = arr.indexOf(videoSrc)\n")
		_, err = f.WriteString("			if (e === '-') {\n")
		_, err = f.WriteString("				counter = arr.indexOf(videoSrc)-1;\n")
		_, err = f.WriteString("			}\n")
		_, err = f.WriteString("			player.src = arr[counter];\n")
		_, err = f.WriteString("			localStorage.setItem(videoSrcTitle, arr[counter]);\n")
		_, err = f.WriteString("		} else {\n")
		_, err = f.WriteString("			player.src = arr[counter];\n")
		_, err = f.WriteString("			localStorage.setItem(videoSrcTitle, arr[counter]);\n")
		_, err = f.WriteString("		}\n")
		_, err = f.WriteString("	}\n")
		_, err = f.WriteString("}\n")
		_, err = f.WriteString("if (counter <= 0 || counter >= arr.length-1) {\n")
		_, err = f.WriteString("	localStorage.removeItem(videoSrcTitle);\n")
		_, err = f.WriteString("}\n")
		_, err = f.WriteString("if ( counter < 0 ) counter += 1;\n")
		_, err = f.WriteString("if ( counter > arr.length-1 ) counter -=1;\n")
		_, err = f.WriteString("player.playbackRate = playbackRate;\n")
		
		_, err = f.WriteString("}\n")
		_, err = f.WriteString("window.addEventListener(\"keydown\", function(event){\n")
		_, err = f.WriteString("    // 1 for playbackRate = 1.0\n")
		_, err = f.WriteString("    if (event.keyCode == 49) {\n")
		_, err = f.WriteString("        playbackRate = 1.0;\n")
		_, err = f.WriteString("        player.playbackRate = playbackRate;\n")
		_, err = f.WriteString("    }\n")
		_, err = f.WriteString("    // 2 for playbackRate = 1.75\n")
		_, err = f.WriteString("    if (event.keyCode == 50) {\n")
		_, err = f.WriteString("        playbackRate = 1.75;\n")
		_, err = f.WriteString("        player.playbackRate = playbackRate;\n")
		_, err = f.WriteString("    }\n")
		_, err = f.WriteString("    // 3 for playbackRate = 3.0\n")
		_, err = f.WriteString("    if (event.keyCode == 51) {\n")
		_, err = f.WriteString("        playbackRate = 3;\n")
		_, err = f.WriteString("        player.playbackRate = playbackRate;\n")
		_, err = f.WriteString("    }\n")
		_, err = f.WriteString("    // 4 - display which video is playing right now in the console\n")
		_, err = f.WriteString("    if (event.keyCode == 52) {\n")
		_, err = f.WriteString("        var a = player.src;\n")
		_, err = f.WriteString("        console.log(a.slice(a.lastIndexOf('/'),a.length));\n")
		_, err = f.WriteString("    }\n")
		_, err = f.WriteString("    // backspace get back \n")
		_, err = f.WriteString("    if (event.keyCode == 8) {\n")
		_, err = f.WriteString("        history.back();\n")
		_, err = f.WriteString("    }\n")
		_, err = f.WriteString("    // left arrow\n")
		_, err = f.WriteString("    if (event.keyCode == 37) {\n")
		_, err = f.WriteString("        event.preventDefault();\n")
		_, err = f.WriteString("        document.querySelector('video').currentTime = document.querySelector('video').currentTime - 5;\n")
		_, err = f.WriteString("    }\n")
		_, err = f.WriteString("    // right arrow\n")
		_, err = f.WriteString("    if (event.keyCode == 39) {\n")
		_, err = f.WriteString("        event.preventDefault();\n")
		_, err = f.WriteString("        document.querySelector('video').currentTime = document.querySelector('video').currentTime + 5;\n")
		_, err = f.WriteString("    }\n")
		_, err = f.WriteString("    // 32 = space\n")
		_, err = f.WriteString("    if (event.keyCode == 32) {\n")
		_, err = f.WriteString("        event.preventDefault();\n")
		_, err = f.WriteString("        togglePlay();\n")
		_, err = f.WriteString("    }\n")
		_, err = f.WriteString("    // 78 = n\n")
		_, err = f.WriteString("    if (event.keyCode == 78) {\n")
		_, err = f.WriteString("        myHandler('+');\n")
		_, err = f.WriteString("    }\n")
		_, err = f.WriteString("    // 80 = p\n")
		_, err = f.WriteString("    if (event.keyCode == 80) {\n")
		_, err = f.WriteString("        myHandler('-');\n")
		_, err = f.WriteString("    }\n")
		_, err = f.WriteString("    // 70 = f\n")
		_, err = f.WriteString("    if (event.keyCode == 70) {\n")
		_, err = f.WriteString("        if (document.fullscreenElement === null) {\n")
		_, err = f.WriteString("	        document.querySelector('video').requestFullscreen();\n")
		_, err = f.WriteString("        } else {\n")
		_, err = f.WriteString("	        document.exitFullscreen();\n")
		_, err = f.WriteString("        }\n")
		_, err = f.WriteString("    }\n")
		_, err = f.WriteString("    // 34 = PageDown\n")
		_, err = f.WriteString("    if (event.keyCode == 34) {\n")
		_, err = f.WriteString("        event.preventDefault();\n")
		_, err = f.WriteString("        playbackRate -= 0.25;\n")
		_, err = f.WriteString("        player.playbackRate = playbackRate;\n")
		_, err = f.WriteString("    }\n")
		_, err = f.WriteString("    // 33 = PageUp\n")
		_, err = f.WriteString("    if (event.keyCode == 33) {\n")
		_, err = f.WriteString("        event.preventDefault();\n")
		_, err = f.WriteString("        playbackRate += 0.25;\n")
		_, err = f.WriteString("        player.playbackRate = playbackRate;\n")
		_, err = f.WriteString("    }\n")
		_, err = f.WriteString("    // 82 = r (remaining videos)\n")
		_, err = f.WriteString("    if (event.keyCode == 82) {\n")
		_, err = f.WriteString("        event.preventDefault();\n")
		_, err = f.WriteString("        console.log\n")
		_, err = f.WriteString(fmt.Sprintf("        console.log(\"%s\")\n", chapter[len(chapter)-1]))
		_, err = f.WriteString("        for (var i = counter; i < arr.length-0; i++) console.log(\"i=\",i,arr[i])\n")
		_, err = f.WriteString("    }\n")
		_, err = f.WriteString("    // 84 = t (total time in the folder of videos)\n")
		_, err = f.WriteString("    if (event.keyCode == 84) {\n")
		_, err = f.WriteString("        event.preventDefault();\n")
		_, err = f.WriteString("        console.log(\"h=\",")
		_, err = f.WriteString(fmt.Sprintln(h, ");"))
		_, err = f.WriteString("        console.log(\"m=\",")
		_, err = f.WriteString(fmt.Sprintln(m, ");"))
		_, err = f.WriteString("        console.log(\"s=\",")
		_, err = f.WriteString(fmt.Sprintln(s, ");"))
		_, err = f.WriteString("    }\n")
		_, err = f.WriteString("});\n")

		_, err = f.WriteString("</script>\n")
		_, err = f.WriteString("</body>\n")
		_, err = f.WriteString("</html>\n")
		check(err)

	}

	var input string
	fmt.Println("Enter Text Below:")
	fmt.Scanln(&input)
	fmt.Println(input)

}
