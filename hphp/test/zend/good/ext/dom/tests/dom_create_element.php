<?hh
<<__EntryPoint>> function main(): void {
print " 1 DOMDocument::createElement('valid')\n";
try {
    $dom = new DOMDocument;
    $dom->createElement('valid');
    print "valid\n";
} catch (Exception $e) {
    print $e->getMessage() . "\n";
}

print " 2 DOMDocument::createElement('-invalid')\n";
try {
    $dom = new DOMDocument;
    $dom->createElement('-invalid');
    print "valid\n";
} catch (Exception $e) {
    print $e->getMessage() . "\n";
}

print " 3 DOMDocument::createElement(' ')\n";
try {
    $dom = new DOMDocument;
    $dom->createElement(' ');
    print "valid\n";
} catch (Exception $e) {
    print $e->getMessage() . "\n";
}

print " 4 DOMDocument::createElement('prefix:valid')\n";
try {
    $dom = new DOMDocument;
    $dom->createElement('prefix:valid');
    print "valid\n";
} catch (Exception $e) {
    print $e->getMessage() . "\n";
}

print " 5 DOMDocument::createElementNS('http://valid.com', 'valid')\n";
try {
    $dom = new DOMDocument;
    $dom->createElementNS('http://valid.com', 'valid');
    print "valid\n";
} catch (Exception $e) {
    print $e->getMessage() . "\n";
}

print " 6 DOMDocument::createElementNS('http://valid.com', 'prefix:valid')\n";
try {
    $dom = new DOMDocument;
    $dom->createElementNS('http://valid.com', 'prefix:valid');
    print "valid\n";
} catch (Exception $e) {
    print $e->getMessage() . "\n";
}

print " 7 DOMDocument::createElementNS('http://valid.com', '-invalid')\n";
try {
    $dom = new DOMDocument;
    $dom->createElementNS('http://valid.com', '-invalid');
    print "valid\n";
} catch (Exception $e) {
    print $e->getMessage() . "\n";
}

print " 8 DOMDocument::createElementNS('http://valid.com', 'prefix:-invalid')\n";
try {
    $dom = new DOMDocument;
    $dom->createElementNS('http://valid.com', 'prefix:-invalid');
    print "valid\n";
} catch (Exception $e) {
    print $e->getMessage() . "\n";
}

print " 9 DOMDocument::createElementNS('', 'prefix:invalid')\n";
try {
    $dom = new DOMDocument;
    $dom->createElementNS('', 'prefix:invalid');
    print "valid\n";
} catch (Exception $e) {
    print $e->getMessage() . "\n";
}

print "10 DOMDocument::createElementNS('http://valid.com', 'prefix:valid:invalid')\n";
try {
    $dom = new DOMDocument;
    $dom->createElementNS('http://valid.com', 'prefix:valid:invalid');
    print "valid\n";
} catch (Exception $e) {
    print $e->getMessage() . "\n";
}

print "11 DOMDocument::createElementNS('http://valid.com', '-prefix:valid')\n";
try {
    $dom = new DOMDocument;
    $dom->createElementNS('http://valid.com', '-prefix:valid');
    print "valid\n";
} catch (Exception $e) {
    print $e->getMessage() . "\n";
}

print "12 DOMDocument::createElementNS('-', 'prefix:valid')\n";
try {
    $dom = new DOMDocument;
    $dom->createElementNS('-', 'prefix:valid');
    print "valid\n";
} catch (Exception $e) {
    print $e->getMessage() . "\n";
}


print "13 DOMElement::__construct('valid')\n";
try {
    $element = new DOMElement('valid');
    print "valid\n";
} catch (Exception $e) {
    print $e->getMessage() . "\n";
}

print "14 DOMElement::__construct('-invalid')\n";
try {
    $element = new DOMElement('-invalid');
    print "valid\n";
} catch (Exception $e) {
    print $e->getMessage() . "\n";
}

print "15 DOMElement::__construct(' ')\n";
try {
    $element = new DOMElement(' ');
    print "valid\n";
} catch (Exception $e) {
    print $e->getMessage() . "\n";
}

print "16 DOMElement::__construct('prefix:valid')\n";
try {
    $element = new DOMElement('prefix:valid');
    print "valid\n";
} catch (Exception $e) {
    print $e->getMessage() . "\n";
}

print "17 DOMElement::__construct('valid', '', 'http://valid.com')\n";
try {
    $element = new DOMElement('valid', '', 'http://valid.com');
    print "valid\n";
} catch (Exception $e) {
    print $e->getMessage() . "\n";
}

print "18 DOMElement::__construct('prefix:valid', '', 'http://valid.com')\n";
try {
    $element = new DOMElement('prefix:valid', '', 'http://valid.com');
    print "valid\n";
} catch (Exception $e) {
    print $e->getMessage() . "\n";
}

print "19 DOMElement::__construct('-invalid', '', 'http://valid.com')\n";
try {
    $element = new DOMElement('-invalid', '', 'http://valid.com');
    print "valid\n";
} catch (Exception $e) {
    print $e->getMessage() . "\n";
}

print "20 DOMElement::__construct('prefix:-invalid', '', 'http://valid.com')\n";
try {
    $element = new DOMElement('prefix:-invalid', '', 'http://valid.com');
    print "valid\n";
} catch (Exception $e) {
    print $e->getMessage() . "\n";
}

print "21 DOMElement::__construct('prefix:invalid', '', '')\n";
try {
    $element = new DOMElement('prefix:invalid', '', '');
    print "valid\n";
} catch (Exception $e) {
    print $e->getMessage() . "\n";
}

print "22 DOMElement::__construct('prefix:valid:invalid', '', 'http://valid.com')\n";
try {
    $element = new DOMElement('prefix:valid:invalid', '', 'http://valid.com');
    print "valid\n";
} catch (Exception $e) {
    print $e->getMessage() . "\n";
}

print "23 DOMElement::__construct('-prefix:valid', '', 'http://valid.com')\n";
try {
    $element = new DOMElement('-prefix:valid', '', 'http://valid.com');
    print "valid\n";
} catch (Exception $e) {
    print $e->getMessage() . "\n";
}

print "24 DOMElement::__construct('prefix:valid', '', '-')\n";
try {
    $element = new DOMElement('prefix:valid', '', '-');
    print "valid\n";
} catch (Exception $e) {
    print $e->getMessage() . "\n";
}

/* the qualifiedName has a prefix and the  namespaceURI is null */

print "25 DOMDocument::createElementNS('', 'prefix:valid')\n";
try {
    $dom = new DOMDocument;
    $dom->createElementNS('', 'prefix:valid');
    print "valid\n";
} catch (Exception $e) {
    print $e->getMessage() . "\n";
}

/* the qualifiedName has a prefix that is "xml" and the  namespaceURI 
   is different from "http://www.w3.org/XML/1998/namespace" [XML Namespaces] */

print "26 DOMDocument::createElementNS('http://wrong.namespaceURI.com', 'xml:valid')\n";
try {
    $dom = new DOMDocument;
    $dom->createElementNS('http://wrong.namespaceURI.com', 'xml:valid');
    print "valid\n";
} catch (Exception $e) {
    print $e->getMessage() . "\n";
}

print "27 DOMElement::__construct('xml:valid', '', 'http://wrong.namespaceURI.com')\n";
try {
    $element = new DOMElement('xml:valid', '', 'http://wrong.namespaceURI.com');
    print "valid\n";
} catch (Exception $e) {
    print $e->getMessage() . "\n";
}

/* This is okay because we reuse the xml namespace from the document */
print "28 DOMDocument::createElementNS('http://www.w3.org/XML/1998/namespace', 'xml:valid')\n";
try {
    $dom = new DOMDocument;
    $dom->createElementNS('http://www.w3.org/XML/1998/namespace', 'xml:valid');
    print "valid\n";
} catch (Exception $e) {
    print $e->getMessage() . "\n";
}

/* This is also valid, as newer libxmls let you create the `xml` namespace */
print "29 DOMElement::__construct('xml:valid', '', 'http://www.w3.org/XML/1998/namespace')\n";
try {
    $element = new DOMElement('xml:valid', '', 'http://www.w3.org/XML/1998/namespace');
    print "valid\n";
} catch (Exception $e) {
    print $e->getMessage() . "\n";
}


/* the qualifiedName or its prefix is "xmlns" and the  namespaceURI is 
   different from  "http://www.w3.org/2000/xmlns/" */

print "30 DOMDocument::createElementNS('http://wrong.namespaceURI.com', 'xmlns:valid')\n";
try {
    $dom = new DOMDocument;
    $dom->createElementNS('http://wrong.namespaceURI.com', 'xmlns:valid');
    print "valid\n";
} catch (Exception $e) {
    print $e->getMessage() . "\n";
}

print "31 DOMElement::__construct('xmlns:valid', '', 'http://wrong.namespaceURI.com')\n";
try {
    $element = new DOMElement('xmlns:valid', '', 'http://wrong.namespaceURI.com');
    print "valid\n";
} catch (Exception $e) {
    print $e->getMessage() . "\n";
}

print "32 DOMDocument::createElementNS('http://www.w3.org/2000/xmlns/', 'xmlns:valid')\n";
try {
    $dom = new DOMDocument;
    $dom->createElementNS('http://www.w3.org/2000/xmlns/', 'xmlns:valid');
    print "valid\n";
} catch (Exception $e) {
    print $e->getMessage() . "\n";
}

print "33 DOMElement::__construct('xmlns:valid', '', 'http://www.w3.org/2000/xmlns/')\n";
try {
    $element = new DOMElement('xmlns:valid', '', 'http://www.w3.org/2000/xmlns/');
    print "valid\n";
} catch (Exception $e) {
    print $e->getMessage() . "\n";
}

/* the namespaceURI is "http://www.w3.org/2000/xmlns/" and neither the 
   qualifiedName nor its prefix is "xmlns". */

print "34 DOMDocument::createElementNS('http://www.w3.org/2000/xmlns/', 'wrongprefix:valid')\n";
try {
    $dom = new DOMDocument;
    $dom->createElementNS('http://www.w3.org/2000/xmlns/', 'wrongprefix:valid');
    print "valid\n";
} catch (Exception $e) {
    print $e->getMessage() . "\n";
}

print "35 DOMElement::__construct('wrongprefix:valid', '', 'http://www.w3.org/2000/xmlns/')\n";
try {
    $element = new DOMElement('wrongprefix:valid', '', 'http://www.w3.org/2000/xmlns/');
    print "valid\n";
} catch (Exception $e) {
    print $e->getMessage() . "\n";
}
}
