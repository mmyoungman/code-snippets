// XCBRenderWeirdGradient.cpp, inspired by Casey Muratori's Handmade Hero
// Uses renderWeirdGradient from Handmade Hero
// The image resizes with the window
// Press Q to quit, F to fullscreen

#include <xcb/xcb.h>
#include <xcb/xcb_image.h>
#include <xcb/xcb_keysyms.h>
#include <X11/keysym.h> // For XK_q XK_Escape etc.

#include <stdlib.h> // For malloc
#include <string.h> // For strlen
#include <stdio.h> // For printf

#define WIDTH 800
#define HEIGHT 600
#define BYTES_PER_PIXEL 4

void renderWeirdGradient(uint8_t *imgdata, int width, int height, int offsetX, int offsetY)
{
    int pitch = width*BYTES_PER_PIXEL;
    uint8_t *row = imgdata;

    for(int y = 0; y < height; y++)
    {
        uint32_t *pixel = (uint32_t *)row;
        for(int x = 0; x < width; x++)
        {
            uint8_t blue = (x + offsetX);
            uint8_t green = (y + offsetY);
            uint8_t red = 0;
            uint8_t alpha = 0;

            *pixel++ = ((alpha << 24) | (red << 16) | (green << 8) | blue);
        }
        row += pitch;
    }
}

int main()
{
    xcb_connection_t *conn;
    xcb_screen_t *screen;
    const xcb_setup_t *setup;

    xcb_pixmap_t pmap;
    xcb_gcontext_t gc;

    uint32_t mask;
    uint32_t values[2];

    // Connect to X server
    conn = xcb_connect(NULL, NULL);
    setup = xcb_get_setup(conn);
    screen = xcb_setup_roots_iterator(setup).data;

    // Create window
    mask = XCB_CW_EVENT_MASK | XCB_CW_BACK_PIXMAP;
    values[0] = screen->black_pixel;
    values[1] = XCB_EVENT_MASK_KEY_PRESS | XCB_EVENT_MASK_KEY_RELEASE |
        XCB_EVENT_MASK_BUTTON_PRESS | XCB_EVENT_MASK_BUTTON_RELEASE |
        XCB_EVENT_MASK_ENTER_WINDOW | XCB_EVENT_MASK_LEAVE_WINDOW |
        XCB_EVENT_MASK_POINTER_MOTION | XCB_EVENT_MASK_EXPOSURE;

    xcb_window_t window = xcb_generate_id(conn);
    xcb_create_window(conn, XCB_COPY_FROM_PARENT, window,
                      screen->root, 
                      // These don't seem to make any difference
                      100, 100,
                      WIDTH, HEIGHT,
                      1, XCB_WINDOW_CLASS_INPUT_OUTPUT,
                      screen->root_visual,
                      mask, values);

    // Name window
    char const *winName = "XCB Window - Press Q to quit, F to toggle fullscreen, or resize the window";
    xcb_change_property (conn,
                         XCB_PROP_MODE_REPLACE,
                         window,
                         XCB_ATOM_WM_NAME, XCB_ATOM_STRING,
                         8,
                         strlen(winName),
                         winName);

    // Set window icon -- DOESN'T WORK
    char const *winIcon= "XCB Icon";
    xcb_change_property (conn,
                         XCB_PROP_MODE_REPLACE,
                         window,
                         XCB_ATOM_WM_ICON_NAME,
                         XCB_ATOM_STRING,
                         8,
                         strlen(winIcon),
                         winIcon);

    // Get image format
    xcb_format_t *format = xcb_setup_pixmap_formats(setup);
    xcb_format_t *formatEnd = format + xcb_setup_pixmap_formats_length(setup);

    while(format != formatEnd)
    {
        if((format->scanline_pad == 32) && 
           (format->depth == 24) && 
           (format->bits_per_pixel == 32))
            break;
        format++;
    }

    // Allocate memory for image data
    uint8_t *imagemem = (uint8_t *)malloc(WIDTH*HEIGHT*BYTES_PER_PIXEL);

    // Create image
    xcb_image_t *image = xcb_image_create(WIDTH, HEIGHT,
                                          XCB_IMAGE_FORMAT_Z_PIXMAP, 
                                          format->scanline_pad, 
                                          format->depth, 
                                          format->bits_per_pixel, 
                                          0, 
                                          (xcb_image_order_t)setup->image_byte_order,
                                          XCB_IMAGE_ORDER_LSB_FIRST,
                                          imagemem, 
                                          WIDTH * HEIGHT * BYTES_PER_PIXEL, 
                                          imagemem);

    // Create pixmap
    pmap = xcb_generate_id(conn);
    xcb_create_pixmap(conn, screen->root_depth, pmap, window, image->width, image->height);

    mask = XCB_GC_FOREGROUND | XCB_GC_BACKGROUND;
    values[0] = screen->black_pixel;
    values[1] = screen->white_pixel; 

    gc = xcb_generate_id(conn);
    xcb_create_gc(conn, gc, pmap, mask, values);

    // Show window
    xcb_map_window(conn, window);

    // Event/render loop
    int offsetX = 0;
    int offsetY = 0;

    int running = 1;

    while(running)
    { 
        xcb_generic_event_t *event;

        // xcb_flush should be done before event loop?
        xcb_flush(conn);

        while((event = xcb_poll_for_event(conn)))
        {
            switch(event->response_type)
            {
                case XCB_KEY_PRESS:
                case XCB_KEY_RELEASE:
                    {
                        // If a key is held down, it gives constant down/up messages. Problem?
                        bool keyIsDown = (event->response_type == XCB_KEY_PRESS);
                        xcb_key_press_event_t *e = (xcb_key_press_event_t *)event;
                        xcb_keysym_t keysym = xcb_key_symbols_get_keysym(xcb_key_symbols_alloc(conn), e->detail, 0);

                        if(keysym == XK_q && keyIsDown)
                        {
                            printf("You pressed Q, so we quit!\n");
                            running = 0;
                        }

                        if(keysym == XK_f && keyIsDown)
                        {
                            printf("You pressed F, so we toggle fullscreen!\n");

                            xcb_intern_atom_cookie_t wmState = xcb_intern_atom(conn, 0, strlen("_NET_WM_STATE"), "_NET_WM_STATE");
                            xcb_intern_atom_cookie_t wmStateFS = xcb_intern_atom(conn, 0, strlen("_NET_WM_STATE_FULLSCREEN"), "_NET_WM_STATE_FULLSCREEN");
                            xcb_client_message_event_t cmEvent;
                            cmEvent.response_type = XCB_CLIENT_MESSAGE;
                            xcb_intern_atom_reply_t *wmStateReply = xcb_intern_atom_reply(conn, wmState, NULL);
                            cmEvent.type = wmStateReply->atom;
                            cmEvent.format = 32;
                            cmEvent.window = window;
                            cmEvent.data.data32[0] = 2; // 2 means toggle fullscreen. 1/2 are add/remove fullscreen
                            xcb_intern_atom_reply_t *wmStateFSReply = xcb_intern_atom_reply(conn, wmStateFS, NULL);
                            cmEvent.data.data32[1] = wmStateFSReply->atom;
                            cmEvent.data.data32[2] = XCB_ATOM_NONE;
                            cmEvent.data.data32[3] = 0;
                            cmEvent.data.data32[4] = 0;

                            xcb_send_event(conn, 1, window,
                                           XCB_EVENT_MASK_SUBSTRUCTURE_REDIRECT | XCB_EVENT_MASK_SUBSTRUCTURE_NOTIFY,
                                           (char*)(&cmEvent));
                            free(wmStateReply);
                            free(wmStateFSReply);
                        } 
                        //printf("Keycode: %d, %d\n", e->detail, keyIsDown);
                    }
                case XCB_BUTTON_PRESS:
                case XCB_BUTTON_RELEASE:
                    {
                        bool buttonIsDown = (event->response_type == XCB_BUTTON_PRESS);
                        xcb_button_press_event_t *e = (xcb_button_press_event_t*)event;
                        //printf("Button: %d, %i, x:%d y:%d\n", e->detail, buttonIsDown, e->event_x, e->event_y);
                        break;
                    }
                case XCB_MOTION_NOTIFY:
                    {
                        xcb_motion_notify_event_t *e = (xcb_motion_notify_event_t*)event;
                        //printf("Pointer x:%d y:%d\n", e->event_x, e->event_y);
                        break;
                    }
                case XCB_EXPOSE:
                    {
                        xcb_expose_event_t *e = (xcb_expose_event_t *)event;

                        // Resize image if the window size has changed
                        if((e->width != image->width) || (e->height != image->height))
                        {
                            // image->data == imagemem, so no free(imagemem) necessary?
                            xcb_image_destroy(image);
                            imagemem = (uint8_t *)malloc(e->width*e->height*BYTES_PER_PIXEL);
                            image = xcb_image_create(e->width, e->height,
                                                     XCB_IMAGE_FORMAT_Z_PIXMAP, 
                                                     format->scanline_pad, 
                                                     format->depth, 
                                                     format->bits_per_pixel, 
                                                     0, 
                                                     (xcb_image_order_t)setup->image_byte_order,
                                                     XCB_IMAGE_ORDER_LSB_FIRST,
                                                     imagemem, 
                                                     e->width * e->height * BYTES_PER_PIXEL, 
                                                     imagemem);
                            xcb_free_pixmap(conn, pmap);
                            xcb_create_pixmap(conn, screen->root_depth, pmap, window, image->width, image->height);
                        }  

                        //printf("Window x:%d y:%d width:%d height:%d\n", e->x, e->y, e->width, e->height);
                        break;
                    }
            }
        }

        free(event);

        renderWeirdGradient(image->data, image->width, image->height, offsetX, offsetY);

        xcb_image_put(conn, pmap, gc, image, 0, 0, 0);
        xcb_copy_area(conn, pmap, window, gc, 0, 0, 0, 0, image->width, image->height);

        offsetX++;
        offsetY += 2;
    }

    xcb_disconnect(conn);
    return 0;
}
