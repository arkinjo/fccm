set term pdf enhanced

set grid

gau(x,m) = (1.0/sqrt(2*pi))*exp(-0.5*(x-m)**2)
rho(x) = (1000*gau(x,-5) + gau(x, 5))/gau(10,5)
set xrange [-10:10]
set out "twoclust.pdf"
set log y
set xlabel "x"
set ylabel "density (arbitrarily scaled)"
set sample 1000
plot rho(x) t ''

unset log y
set ylabel "Membership function"
set yrange [-0.1:1.18]

set out "egfcm.pdf"
p "< grep Data0 example.dat | sort -n -k 3" u 3:4 t 'u_1(x)' w l \
 ,'' u 3:5 t 'u_2(x)' w l

set out "egfcc.pdf"
p "< grep Data1 example.dat | sort -n -k 3" u 3:4 t 'u_1(x)' w l \
 ,'' u 3:5 t 'u_2(x)' w l

