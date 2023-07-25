<?php

class VowelConsonantProcessor {
    private $vowels = ['a', 'i', 'u', 'e', 'o'];

    private function extractCharacters(string $param, bool $isVowel): array {
        $characters = [];

        $char = strtolower(str_replace(" ", "", $param));

        for ($i = 0; $i < strlen($char); $i++) {
            $isVowelChar = in_array($char[$i], $this->vowels);

            if ($isVowel === $isVowelChar) {
                $characters[] = $char[$i];
            }
        }

        return $characters;
    }


    public function procVowel(string $input):string {
        $vowels = $this->extractCharacters($input, true);
        sort($vowels);
        return implode('', $vowels);
    }

    public function procConsonant(string $input):string {
        return implode('',$this->extractCharacters($input, false));
    }
}

// Main function
echo "Input one line of words (S) : ";
$input = trim(fgets(STDIN));

$processor = new VowelConsonantProcessor();
$charVowel = $processor->procVowel($input);
$charConsonant = $processor->procConsonant($input);

echo "Vowel Characters :\n";
echo $charVowel . "\n";
echo "Consonant Characters :\n";
echo $charConsonant . "\n";

?>