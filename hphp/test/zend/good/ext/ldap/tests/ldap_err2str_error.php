<?hh
<<__EntryPoint>> function main(): void {
// Too few args
try { var_dump(ldap_err2str()); } catch (Exception $e) { echo "\n".'Warning: '.$e->getMessage().' in '.__FILE__.' on line '.__LINE__."\n"; }

// Too many args
try { var_dump(ldap_err2str(1, "Additional data")); } catch (Exception $e) { echo "\n".'Warning: '.$e->getMessage().' in '.__FILE__.' on line '.__LINE__."\n"; }
echo "===DONE===\n";
}
