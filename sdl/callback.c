/*
  SDL Go Wrapper
  Copyright (C) 2013 Kristoffer Gronlund <kgronlund@suse.com>

  This software is provided 'as-is', without any express or implied
  warranty.  In no event will the authors be held liable for any damages
  arising from the use of this software.

  Permission is granted to anyone to use this software for any purpose,
  including commercial applications, and to alter it and redistribute it
  freely, subject to the following restrictions:

  1. The origin of this software must not be misrepresented; you must not
     claim that you wrote the original software. If you use this software
     in a product, an acknowledgment in the product documentation would be
     appreciated but is not required.
  2. Altered source versions must be plainly marked as such, and must not be
     misrepresented as being the original software.
  3. This notice may not be removed or altered from any source distribution.
*/

#include <SDL2/SDL_audio.h>

extern void go_sdl2_audio_callback(void* userdata, Uint8* stream, int len);

SDL_AudioCallback go_sdl2_get_callback() {
	return &go_sdl2_audio_callback;
}
