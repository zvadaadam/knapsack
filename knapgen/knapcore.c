#include "knapcore.h"
#include <math.h>
#include <stdlib.h>

/*-----------------------------------------------------------------------------
	flat distribution over the range low .. high inclusive
-----------------------------------------------------------------------------*/
static int rngrandom (int low, int high) {
    int q = high-low+1;
    return random()%q + low;
}
/*-----------------------------------------------------------------------------

-----------------------------------------------------------------------------*/
static int go (
    int w,			/* weight, 1..wm */
    int wm,			/* max weight */
    double k,			/* exponent */
    int dir			/* -1..more small things, 1..more big things, 0..does not mattter */
) {
    double thr;
    switch (dir) {
        default:
        case  0: return 1;
        case -1: thr = RAND_MAX / pow (w,k); break;
        case  1: thr = RAND_MAX / pow (wm-w+1, k); break;
    }
    if (thr >= random()) return 1;
    return 0;
}

/*-----------------------------------------------------------------------------*/
int knapcore (		/* returns total weight */
	int* weights,
	int* costs,
	int  n,		/* total things # */
	int  Wmax,	/* max generated weight */
	int  Cmax,	/* max cost (min cost=1) */
	double ke,	/* exponent */
	int  d		/* -1..more small things, 1..more big things, 0..does not mattter */
) {

    int*   issued=(int*)malloc((Wmax+1)*sizeof(int));
    int  k; 
    long tw, w; 

    if (!issued) return 0;	/* indicates memory alloc failure */
    
    for (k=0; k<Wmax; k++) issued[k]=0;
    tw=0;

    k=0;
    while(k<n) {
        w = rngrandom(1,Wmax);
        if (issued[w] > 0) continue;	/* already have a thing with this weight */
        if (go(w, Wmax, ke, d)) {
            costs[k] = issued[w] = rngrandom(1,Cmax);
            weights[k] = w;
            tw += w;
            k++;
        }
    }
    free (issued);
    return tw;
}

