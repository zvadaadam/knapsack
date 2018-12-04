#ifndef KNAPCORE_H
#define KNAPCORE_H
extern
int knapcore (		/* returns total weight */
	int* weights,
	int* costs,
	int  n,		/* total things # */
	int  Wmax,	/* max generated weight */
	int  Cmax,	/* max cost (min cost=1) */
	double ke,	/* exponent */
	int  d		/* -1..more small things, 1..more big things, 0..does not mattter */
);
#endif
