#include <stdio.h>
#include <stdlib.h>

int main(void)
{
  int i;

  printf("Gerando 10 valores aleatorios:\n\n");
  
  for (i = 0; i < 10; i++)
  {
    /* gerando valores aleatórios entre zero e 100 */
    printf("%d ", rand() % 100);
  }
  
  return 0;
}