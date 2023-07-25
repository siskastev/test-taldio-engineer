<?php

// max passengers = 4

// input n = 8 => 2 3 4 4 2 1 3 1

// 2+3 = 5
// 5 >= 4 true
//c = 5-4 = 1 got 1 true

//1+4 = 5
// 5 >= 4 true
//c = 5-4 = 1 got 1 true

//c = 1+4 = 5
// 5 >= 4 true
//c = 5-4 = 1 got 1 true

//c=1+2 = 3 when 3< 4 is true so continue loop
//c=3+1 = 4 dan === 4 got 1 true
//c= 4-4 =0

//c= 0 + 3 = 3
// c= 3+1 = 4 dan === 4 got 1 true
//c= 4-4 =0

// so total = 5

function calculateMinBuses(array $familyMembers): int {
    $maxPassengers = 4;
    $totalBuses = 0;
    $currentPassengers = 0;

    foreach ($familyMembers as $members) {
        $currentPassengers += $members;
        while ($currentPassengers >= $maxPassengers) {
            $totalBuses++;
            $currentPassengers -= $maxPassengers;
        }
    }

    if ($currentPassengers > 0) {
        $totalBuses++;
    }

    return $totalBuses;
}

echo "Input the number of families : ";
$n = (int) trim(fgets(STDIN));

echo "Input the number of members in each family (separated by a space): ";
$inputMembers = trim(fgets(STDIN));

$familyMembers = explode(' ', $inputMembers);
$familyMembers = array_map('intval', $familyMembers);

if (count($familyMembers) !== $n) {
    echo "Error: ​Input must be equal with count of family.\n";
    exit();
}

echo "Minimum bus required is : ".calculateMinBuses($familyMembers);
