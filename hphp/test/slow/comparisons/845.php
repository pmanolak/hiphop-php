<?hh



<<__EntryPoint>>
function main_845() :mixed{
$i = 0;
 ++$i; print $i;
 print "\t";
 print (false==true) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = false;
 print ($a ==true) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = true;
 print (false==$b) ? 'Y' : 'N';
 print ($a ==$b) ? 'Y' : 'N';
 print "\t";
 print "false == true	";
 print "\n";
 ++$i; print $i;
 print "\t";
 print (false==false) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = false;
 print ($a ==false) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = false;
 print (false==$b) ? 'Y' : 'N';
 print ($a ==$b) ? 'Y' : 'N';
 print "\t";
 print "false == false	";
 print "\n";
 ++$i; print $i;
 print "\t";
 print (false==1) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = false;
 print ($a ==1) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = 1;
 print (false==$b) ? 'Y' : 'N';
 print ($a ==$b) ? 'Y' : 'N';
 print "\t";
 print "false == 1	";
 print "\n";
 ++$i; print $i;
 print "\t";
 print (HH\Lib\Legacy_FIXME\eq(false, 0)) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = false;
 print (HH\Lib\Legacy_FIXME\eq($a, 0)) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = 0;
 print (HH\Lib\Legacy_FIXME\eq(false, $b)) ? 'Y' : 'N';
 print (HH\Lib\Legacy_FIXME\eq($a, $b)) ? 'Y' : 'N';
 print "\t";
 print "false == 0	";
 print "\n";
 ++$i; print $i;
 print "\t";
 print (false==-1) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = false;
 print ($a ==-1) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = -1;
 print (false==$b) ? 'Y' : 'N';
 print ($a ==$b) ? 'Y' : 'N';
 print "\t";
 print "false == -1	";
 print "\n";
 ++$i; print $i;
 print "\t";
 print (false=='1') ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = false;
 print ($a =='1') ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = '1';
 print (false==$b) ? 'Y' : 'N';
 print ($a ==$b) ? 'Y' : 'N';
 print "\t";
 print "false == '1'	";
 print "\n";
 ++$i; print $i;
 print "\t";
 print (HH\Lib\Legacy_FIXME\eq(false, '0')) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = false;
 print (HH\Lib\Legacy_FIXME\eq($a, '0')) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = '0';
 print (HH\Lib\Legacy_FIXME\eq(false, $b)) ? 'Y' : 'N';
 print (HH\Lib\Legacy_FIXME\eq($a, $b)) ? 'Y' : 'N';
 print "\t";
 print "false == '0'	";
 print "\n";
 ++$i; print $i;
 print "\t";
 print (false=='-1') ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = false;
 print ($a =='-1') ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = '-1';
 print (false==$b) ? 'Y' : 'N';
 print ($a ==$b) ? 'Y' : 'N';
 print "\t";
 print "false == '-1'	";
 print "\n";
 ++$i; print $i;
 print "\t";
 print (HH\Lib\Legacy_FIXME\eq(false, null)) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = false;
 print (HH\Lib\Legacy_FIXME\eq($a, null)) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = null;
 print (HH\Lib\Legacy_FIXME\eq(false, $b)) ? 'Y' : 'N';
 print (HH\Lib\Legacy_FIXME\eq($a, $b)) ? 'Y' : 'N';
 print "\t";
 print "false == null	";
 print "\n";
 ++$i; print $i;
 print "\t";
 print (HH\Lib\Legacy_FIXME\eq(false, dict[])) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = false;
 print (HH\Lib\Legacy_FIXME\eq($a, dict[])) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = dict[];
 print (HH\Lib\Legacy_FIXME\eq(false, $b)) ? 'Y' : 'N';
 print (HH\Lib\Legacy_FIXME\eq($a, $b)) ? 'Y' : 'N';
 print "\t";
 print "false == array()	";
 print "\n";
 ++$i; print $i;
 print "\t";
 print (false==vec[1]) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = false;
 print ($a ==vec[1]) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = vec[1];
 print (false==$b) ? 'Y' : 'N';
 print ($a ==$b) ? 'Y' : 'N';
 print "\t";
 print "false == array(1)	";
 print "\n";
 ++$i; print $i;
 print "\t";
 print (false==vec[2]) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = false;
 print ($a ==vec[2]) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = vec[2];
 print (false==$b) ? 'Y' : 'N';
 print ($a ==$b) ? 'Y' : 'N';
 print "\t";
 print "false == array(2)	";
 print "\n";
 ++$i; print $i;
 print "\t";
 print (false==vec['1']) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = false;
 print ($a ==vec['1']) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = vec['1'];
 print (false==$b) ? 'Y' : 'N';
 print ($a ==$b) ? 'Y' : 'N';
 print "\t";
 print "false == array('1')	";
 print "\n";
 ++$i; print $i;
 print "\t";
 print (false==dict['0' => '1']) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = false;
 print ($a ==dict['0' => '1']) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = dict['0' => '1'];
 print (false==$b) ? 'Y' : 'N';
 print ($a ==$b) ? 'Y' : 'N';
 print "\t";
 print "false == array('0' => '1')	";
 print "\n";
 ++$i; print $i;
 print "\t";
 print (false==vec['a']) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = false;
 print ($a ==vec['a']) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = vec['a'];
 print (false==$b) ? 'Y' : 'N';
 print ($a ==$b) ? 'Y' : 'N';
 print "\t";
 print "false == array('a')	";
 print "\n";
 ++$i; print $i;
 print "\t";
 print (false==dict['a' => 1]) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = false;
 print ($a ==dict['a' => 1]) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = dict['a' => 1];
 print (false==$b) ? 'Y' : 'N';
 print ($a ==$b) ? 'Y' : 'N';
 print "\t";
 print "false == array('a' => 1)	";
 print "\n";
 ++$i; print $i;
 print "\t";
 print (false==dict['b' => 1]) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = false;
 print ($a ==dict['b' => 1]) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = dict['b' => 1];
 print (false==$b) ? 'Y' : 'N';
 print ($a ==$b) ? 'Y' : 'N';
 print "\t";
 print "false == array('b' => 1)	";
 print "\n";
 ++$i; print $i;
 print "\t";
 print (false==dict['a' => 1, 'b' => 2]) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = false;
 print ($a ==dict['a' => 1, 'b' => 2]) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = dict['a' => 1, 'b' => 2];
 print (false==$b) ? 'Y' : 'N';
 print ($a ==$b) ? 'Y' : 'N';
 print "\t";
 print "false == array('a' => 1, 'b' => 2)	";
 print "\n";
 ++$i; print $i;
 print "\t";
 print (false==vec[dict['a' => 1]]) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = false;
 print ($a ==vec[dict['a' => 1]]) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = vec[dict['a' => 1]];
 print (false==$b) ? 'Y' : 'N';
 print ($a ==$b) ? 'Y' : 'N';
 print "\t";
 print "false == array(array('a' => 1))	";
 print "\n";
 ++$i; print $i;
 print "\t";
 print (false==vec[dict['b' => 1]]) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = false;
 print ($a ==vec[dict['b' => 1]]) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = vec[dict['b' => 1]];
 print (false==$b) ? 'Y' : 'N';
 print ($a ==$b) ? 'Y' : 'N';
 print "\t";
 print "false == array(array('b' => 1))	";
 print "\n";
 ++$i; print $i;
 print "\t";
 print (false=='php') ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = false;
 print ($a =='php') ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = 'php';
 print (false==$b) ? 'Y' : 'N';
 print ($a ==$b) ? 'Y' : 'N';
 print "\t";
 print "false == 'php'	";
 print "\n";
 ++$i; print $i;
 print "\t";
 print (HH\Lib\Legacy_FIXME\eq(false, '')) ? 'Y' : 'N';
 $a = 1;
 $a = 't';
 $a = false;
 print (HH\Lib\Legacy_FIXME\eq($a, '')) ? 'Y' : 'N';
 $b = 1;
 $b = 't';
 $b = '';
 print (HH\Lib\Legacy_FIXME\eq(false, $b)) ? 'Y' : 'N';
 print (HH\Lib\Legacy_FIXME\eq($a, $b)) ? 'Y' : 'N';
 print "\t";
 print "false == ''	";
 print "\n";
}
