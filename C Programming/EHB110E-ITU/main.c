/*
Name Surname | StudentID
Ulaş Erdoğan | 040190230
*/

#include <stdio.h>
#include "bitmap.h"

void blurBMP(); // blurring operation
unsigned char createSmoothedPixel(unsigned char *framed, int i, int j, int rgb, int width); // funciton that applies blurring to pixels
void downscaleBMP(); // downscaling operation

int main()
{
    blurBMP();      // blurring operation
    downscaleBMP(); // downscaling operation

    return 0;
}

void blurBMP()
{
    // open bitmap file to read
    FILE *bitmap;
    bitmap = fopen("itu.bmp", "rb");

    // check the condition of the file and warn on the errors
    if (bitmap == NULL)
    {
        printf("Bitmap file could not open! Blurring failed...\n");
    }
    else
    {
        printf("Bitmap file have started to blurring opearation!\n");
    }

    // place the header struct of the main bitmap file to structs
    BITMAPFILEHEADER fileHeader;
    BITMAPINFOHEADER infoHeader;
    fread(&fileHeader, sizeof(fileHeader), 1, bitmap); // file header
    fread(&infoHeader, sizeof(infoHeader), 1, bitmap); // info header

    // creating the array to keep the bitmaps pixel data and reading
    unsigned char *pixels = (unsigned char *)malloc(fileHeader.bfSize - 54);
    fread(pixels, fileHeader.bfSize - 54, 1, bitmap);
    // closing the read bitmap
    fclose(bitmap);

    // defining an arrray to keep the framed pixel data of the bitmap
    unsigned char *framed_pixels = (unsigned char *)malloc((infoHeader.biWidth + 2) * (infoHeader.biHeight + 2) * 3);

    // the black framing operation
    for (int i = 0; i < infoHeader.biHeight + 2; i++)
    {
        for (int j = 0; j < infoHeader.biWidth + 2; j++)
        {
            // if it is on the edges of the photo, it writes 0 rgb values for black frame
            if (i == 0 || j == 0 || i == (infoHeader.biHeight + 1) || j == (infoHeader.biWidth + 1))
            {
                framed_pixels[i * 3 * (infoHeader.biWidth + 2) + (j * 3)] = 0;
                framed_pixels[i * 3 * (infoHeader.biWidth + 2) + (j * 3) + 1] = 0;
                framed_pixels[i * 3 * (infoHeader.biWidth + 2) + (j * 3) + 2] = 0;
            }
            // if it is not on the edgesi it writes the data from the original pixel data
            else
            {
                framed_pixels[i * 3 * (infoHeader.biWidth + 2) + (j * 3)] = pixels[(i - 1) * 3 * (infoHeader.biWidth) + ((j - 1) * 3)];
                framed_pixels[i * 3 * (infoHeader.biWidth + 2) + (j * 3) + 1] = pixels[(i - 1) * 3 * (infoHeader.biWidth) + ((j - 1) * 3) + 1];
                framed_pixels[i * 3 * (infoHeader.biWidth + 2) + (j * 3) + 2] = pixels[(i - 1) * 3 * (infoHeader.biWidth) + ((j - 1) * 3) + 2];
            }
        }
    }

    free(pixels); // unloading allocated ram, which won't be used again

    // defining an arrray to keep the wanted form of the bitmap
    unsigned char *blurred_pixels = (unsigned char *)malloc(fileHeader.bfSize - 54);

    // it processes dot product for every pixel
    for (int i = 0; i < infoHeader.biHeight; i++) // loop for columns
    {
        for (int j = 0; j < infoHeader.biWidth; j++) // loop for rows
        {
            blurred_pixels[i * 3 * infoHeader.biWidth + j * 3] = createSmoothedPixel(framed_pixels, i, j, 0, infoHeader.biWidth);     // process for r value
            blurred_pixels[i * 3 * infoHeader.biWidth + j * 3 + 1] = createSmoothedPixel(framed_pixels, i, j, 1, infoHeader.biWidth); // process for g value
            blurred_pixels[i * 3 * infoHeader.biWidth + j * 3 + 2] = createSmoothedPixel(framed_pixels, i, j, 2, infoHeader.biWidth); // process for b value
        }
    }

    free(framed_pixels); // unloading allocated ram, which won't be used again

    // create a bitmap file to write the blurred version
    FILE *blurred_bitmap;
    blurred_bitmap = fopen("blurred_itu.bmp", "wb");

    // write the header of the blurred version
    fwrite(&fileHeader, sizeof(fileHeader), 1, blurred_bitmap); // blurred file header
    fwrite(&infoHeader, sizeof(infoHeader), 1, blurred_bitmap); // blurred info header

    // write the pixel data of the blurred version
    fwrite(blurred_pixels, fileHeader.bfSize - 54, 1, blurred_bitmap);

    free(blurred_pixels);   // unloading allocated ram, which won't be used again
    fclose(blurred_bitmap); // closing the processed photo
}

// a function to read 3x3 pixel data which will be applied dot product and give the 1 pixel result by dot product
// it gets arguement as the framed bitmap pixels, the row and column (i-j), an indicator to understand which values it will detect (0 for r, 1 for g, 2 for b: how bytes far from the beginning of the pixel) and the width of the bitmap
unsigned char createSmoothedPixel(unsigned char *framed, int i, int j, int rgb, int width)
{
    int pixelData = 0; //keeps the result of the dot product

    // dot product resulting
    pixelData += framed[i * 3 * (width + 2) + j * 3 + rgb] * 0.0625;
    pixelData += framed[i * 3 * (width + 2) + j * 3 + 3 + rgb] * 0.125;
    pixelData += framed[i * 3 * (width + 2) + j * 3 + 6 + rgb] * 0.0625;
    pixelData += framed[(i + 1) * 3 * (width + 2) + j * 3 + rgb] * 0.125;
    pixelData += framed[(i + 1) * 3 * (width + 2) + j * 3 + 3 + rgb] * 0.25;
    pixelData += framed[(i + 1) * 3 * (width + 2) + j * 3 + 6 + rgb] * 0.125;
    pixelData += framed[(i + 2) * 3 * (width + 2) + j * 3 + rgb] * 0.0625;
    pixelData += framed[(i + 2) * 3 * (width + 2) + j * 3 + 3 + rgb] * 0.125;
    pixelData += framed[(i + 2) * 3 * (width + 2) + j * 3 + 6 + rgb] * 0.0625;

    // it returns processed pixels r, g or b value
    return pixelData;
}

void downscaleBMP()
{
    // open bitmap file to read
    FILE *bitmap;
    bitmap = fopen("itu.bmp", "rb");

    // check the condition of the file and warn on the errors
    if (bitmap == NULL)
    {
        printf("Bitmap file could not open! Downscaling failed...\n");
    }
    else
    {
        printf("Bitmap file have started to downscaling operation!\n");
    }

    // place the header struct of the main bitmap file to structs
    BITMAPFILEHEADER fileHeader;
    BITMAPINFOHEADER infoHeader;
    fread(&fileHeader, sizeof(fileHeader), 1, bitmap); // file header
    fread(&infoHeader, sizeof(infoHeader), 1, bitmap); // info header

    // change the values of the headers to set them proper to the downscaled version
    fileHeader.bfSize = (fileHeader.bfSize - 54) / 4 + 54;
    infoHeader.biWidth /= 2;
    infoHeader.biHeight /= 2;

    // create a bitmap file to write the downscaled version
    FILE *downscaled_bitmap;
    downscaled_bitmap = fopen("downscaled_itu.bmp", "wb");

    // write the downscaled versions headers
    fwrite(&fileHeader, sizeof(fileHeader), 1, downscaled_bitmap); // downscaled file header
    fwrite(&infoHeader, sizeof(infoHeader), 1, downscaled_bitmap); // downscaled info header

    // the algorithm works by creating downscaled bitmaps pixels row by row
    // in the first for loop it represents the column pixels of the downscaled bitmap
    for (int i = 0; i < infoHeader.biHeight + 4; i++)
    {
        // in the second for loop it represents the row pixels of the downscaled bitmap
        for (int j = 0; j < infoHeader.biWidth; j++)
        {
            unsigned char r = 0, g = 0, b = 0; // r, g and b variables keeps a pixel of downscaled bitmaps r-g-b values
            unsigned char tr, tg, tb;          // these are temporary values to make addition for the read pixels' r-g-b values

            // move cursor and read 2 pixels' r-g-b data to write the "i"th indexed rows and "j"th indexed columns pixels
            fseek(bitmap, 54 + (i * 2 * 3 * infoHeader.biWidth * 2) + j * 3 * 2, SEEK_SET);
            fread(&tr, sizeof(unsigned char), 1, bitmap);
            r += tr / 4;
            fread(&tg, sizeof(unsigned char), 1, bitmap);
            g += tg / 4;
            fread(&tb, sizeof(unsigned char), 1, bitmap);
            b += tb / 4;
            fread(&tr, sizeof(unsigned char), 1, bitmap);
            r += tr / 4;
            fread(&tg, sizeof(unsigned char), 1, bitmap);
            g += tg / 4;
            fread(&tb, sizeof(unsigned char), 1, bitmap);
            b += tb / 4;

            // move cursor to row below and read 2 pixels' r-g-b data in line to write same row-columned pixel
            fseek(bitmap, 54 + ((i * 2 + 1) * 3 * infoHeader.biWidth * 2) + j * 3 * 2, SEEK_SET);
            fread(&tr, sizeof(unsigned char), 1, bitmap);
            r += tr / 4;
            fread(&tg, sizeof(unsigned char), 1, bitmap);
            g += tg / 4;
            fread(&tb, sizeof(unsigned char), 1, bitmap);
            b += tb / 4;
            fread(&tr, sizeof(unsigned char), 1, bitmap);
            r += tr / 4;
            fread(&tg, sizeof(unsigned char), 1, bitmap);
            g += tg / 4;
            fread(&tb, sizeof(unsigned char), 1, bitmap);
            b += tb / 4;

            // we got quarter values (average of values) of 4 pixels r-g-b's
            // then, we write them to the downscaled bitmap
            fwrite(&r, sizeof(unsigned char), 1, downscaled_bitmap);
            fwrite(&g, sizeof(unsigned char), 1, downscaled_bitmap);
            fwrite(&b, sizeof(unsigned char), 1, downscaled_bitmap);
        }
    }

    // close the opened bitmap files
    fclose(bitmap);
    fclose(downscaled_bitmap);
}