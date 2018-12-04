#include <stdio.h>
#include <stdlib.h>	/* strtol() */
#include <time.h>   /* time()   */
#include "knapcore.h"

/*-----------------------------------------------------------------------------
Synopsis: knapgen <options>
	-I <initial id> defaults to 0
	-n <total things #> 
	-N <instances #>
	-m <the ratio of max knapsack capacity to total weight> 
	-W <max weight>
	-C <max cost>
	-k <exponent (real)>
	-d <-1, 0, 1; -1..more small things, 1..more large things, 0..balance>
-------------------------------------------------------------------------------*/
static int arg (char opt, int argc, char* argv[]) {
    int i, match=0; char* es;
    long val;
    for (i=1; i<argc; i++) {
	if (argv[i][0] == '-' && argv[i][1] == opt && argv[i][2] == '\0') {
	    match=1;
	} else if (match) {
	    val = strtol (argv[i],&es,10);
	    if (*es) {
		fprintf (stderr, "-%c option not integer\n", opt);
		exit(0);
	    }
	    return val;
	}
    }
    fprintf (stderr, "-%c option missing\n", opt);
    exit(0);
}
/*-----------------------------------------------------------------------------*/
static int argdf (char opt, int argc, char* argv[], int df) {
    int i, match=0; char* es;
    long val;
    for (i=1; i<argc; i++) {
	if (argv[i][0] == '-' && argv[i][1] == opt && argv[i][2] == '\0') {
	    match=1;
	} else if (match) {
	    val = strtol (argv[i],&es,10);
	    if (*es) {
		fprintf (stderr, "-%c option not integer\n", opt);
		exit(0);
	    }
	    return val;
	}
    }
    return df;
}
/*-----------------------------------------------------------------------------*/
static int targ (char opt, int argc, char* argv[]) {
    int i, match=0; char* es;
    long val;
    for (i=1; i<argc; i++) {
	if (argv[i][0] == '-' && argv[i][1] == opt && argv[i][2] == '\0') {
	    match=1;
	} else if (match) {
	    val = strtol (argv[i],&es,10);
	    if (*es) {
		fprintf (stderr, "-%c option not integer\n", opt);
		exit(0);
	    }
	    if (val < -1 || val > 1) {
		fprintf (stderr, "-%c option value should be -1 .. 1\n", opt);
		exit(0);
	    }
	    return val;
	}
    }
    fprintf (stderr, "-%c option missing\n", opt);
    exit(0);
}
/*-----------------------------------------------------------------------------*/
static double farg (char opt, int argc, char* argv[]) {
    int i, match=0; char* es;
    double val;
    for (i=1; i<argc; i++) {
	if (argv[i][0] == '-' && argv[i][1] == opt && argv[i][2] == '\0') {
	    match=1;
	} else if (match) {
	    val = strtod (argv[i],&es);
	    if (*es) {
		fprintf (stderr, "-%c option not in float format\n", opt);
		exit(0);
	    }
	    return val;
	}
    }
    fprintf (stderr, "-%c option missing\n", opt);
    exit(0);
}
/*-----------------------------------------------------------------------------*/
int main (int argc, char* argv[]) {

    int    I 	= argdf ('I',argc,argv,0);
    int    n 	= arg   ('n',argc,argv);
    int    N 	= arg   ('N',argc,argv);
    double m 	= farg  ('m',argc,argv);
    double ke 	= farg  ('k',argc,argv);
    int    Wmax = arg   ('W',argc,argv);
    int    Cmax = arg   ('C',argc,argv);
    int    d    = targ  ('d',argc,argv);

    long ttw, tw; int i,j; int M;
    int* weights = (int*)malloc(n*sizeof(int));
    int* costs   = (int*)malloc(n*sizeof(int));

    srandom(time(NULL));
    ttw=0;
    for (j=0; j<N; j++) {
        tw=knapcore (weights, costs, n, Wmax, Cmax, ke, d); ttw+=tw; M=m*tw;
	printf ("%d %d %d",I+j, n, M);
	for (i=0; i<n; i++) printf (" %d %d", weights[i], costs[i]);
	printf ("\n");
        fprintf (stderr, "total weight %ld\n", tw);
    }
    fprintf (stderr, "average total weight %.2f; first unused instance id %d\n",
	(double) ttw / (double) N, I+N);
    return 0;
}

