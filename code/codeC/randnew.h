#include <math.h>
#include <time.h>
#include <stdio.h>

#ifndef _RANDNEW_H
#define _RANDNEW_H

int randnew ();
void init_seed ();

int rand_seed = 10;

int randnew ()
{
   rand_seed = rand_seed * 1103515245 + 12345;
   return (unsigned int) (rand_seed / 65536) % 32768;
}

void init_seed ()
{
   rand_seed = time (NULL);
}

void Print(int i) {
 	printf("%d\n", i);
}

#endif